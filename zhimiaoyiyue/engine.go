package zhimiaoyiyue

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
	"zmyy_seckill/config"
	"zmyy_seckill/consts"
	"zmyy_seckill/model"
)

type ZMYYEngine struct {
	Conf config.CustomerConf
}

type SecKill interface {
	GetCustomerList() (*model.CustomerList, error)
}

var wg2 sync.WaitGroup
var flag bool
var detailLock sync.Mutex

func (e *ZMYYEngine) Init() {
	var c config.RootConf
	conf, err2 := c.GetConf()
	if err2 != nil {
		panic(err2)
	}
	e.Conf = conf
	log.SetPrefix("【zmyy-seckill】")
}

/**
获取疫苗可预约日期
*/
func getAvailableDates(customerId, productId, month int,
	subscribeDateChan chan<- *model.SubscribeDate, e *ZMYYEngine, ctx context.Context, stop context.CancelFunc, ip ...string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			log.Printf("%v 正在获取dates\n", ip)
			subscribeDates := e.GetCustSubscribeDateAll(customerId, productId, month, ip...)
			if subscribeDates == nil || len(subscribeDates.Dates) == 0 {
				log.Printf("%v 未获取到可预约日期,尝试重新获取...\n", ip)
			} else {
				subscribeDateChan <- subscribeDates
				stop()
			}
		}

	}
}
func deleteSliceSafe(availableDates []model.Dates, v model.Dates, index int) []model.Dates {
	detailLock.Lock()
	if index < len(availableDates) && v == availableDates[index] {
		availableDates = append(availableDates[:index], availableDates[index+1:]...)
	} else {
		for i := 0; i < len(availableDates); i++ {
			if v == availableDates[i] {
				availableDates = append(availableDates[:i], availableDates[i+1:]...)
				break
			}
		}
	}
	detailLock.Unlock()
	return availableDates
}

/**
获取筛选后的日期的可预约时间段
*/
func getDateDetail(dateDetailCh chan model.DateDetail,
	productId int, customerId int, e *ZMYYEngine,
	ctx context.Context, stop context.CancelFunc,
	ip ...string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			//如果日期全部请求完毕，则返回
			if len(availableDates) == 0 {
				stop()
				return
			}
			//随机从availableDates切片中取一个日期进行请求,加锁防止其他线程删除切片元素
			detailLock.Lock()
			index := rand.Intn(len(availableDates))
			v := availableDates[index]
			detailLock.Unlock()
			log.Printf("%v 尝试获取 %s 疫苗信息...\n", ip, v.Date)
			dateDetails := e.GetCustSubscribeDateDetail(v.Date, productId, customerId, ip...)
			if dateDetails == nil {
				log.Printf("%v 未找到 %s 的可预约时间，尝试查找下一个时间...\n", ip, v.Date)
			} else {
				//找到日期具体时间段后将日期从切片中删除,需要注意原子性问题，故加锁并二次确认当前位置是否还是该日期值
				availableDates = deleteSliceSafe(availableDates, v, index)
				//处理获取到的具体预约时间段,多线程情况下可能添加重复日期，先不管
				for _, v := range dateDetails.DateDetails {
					if v.Qty <= 0 {
						continue
					}
					v.Date = dateDetails.Date
					dateDetailCh <- v
				}
			}
		}
	}

}

var availableDates = make([]model.Dates, 0)

func (e *ZMYYEngine) Run(customerId, productId int, startTime time.Time) {
	var subscribeDateCh = make(chan *model.SubscribeDate)
	var dateDetailCh = make(chan model.DateDetail, 50)
	//1. 使用多个ip多线程获取疫苗可预约的日期
	getDatesCtx, stopGetDates := context.WithCancel(context.Background())
	for i := 0; i < len(consts.ProxyIpArr); i++ {
		go getAvailableDates(customerId, productId, e.Conf.Month, subscribeDateCh,
			e, getDatesCtx, stopGetDates, consts.ProxyIpArr[i])
	}
	log.Println("正在等待获取日期...")
	subscribeDates := <-subscribeDateCh
	//2.将获取到的日期进行筛选，因为有的日期已经被约满了
	cnt := 0
	for _, v := range subscribeDates.Dates {
		if v.Enable == false {
			continue
		}
		cnt++
		fmt.Printf("日期：%s \n", v.Date)
		availableDates = append(availableDates, v)
	}
	fmt.Printf("共找到 %d个 可预约的日期\n", cnt)
	log.Println("正在等待获取预约的时间段...")
	//3.获取筛选后的日期的可预约时间段
	getDetailCtx, stopGetDetail := context.WithCancel(context.Background())
	for i := 0; i < len(consts.ProxyIpArr); i++ {
		go getDateDetail(dateDetailCh, productId,
			customerId, e, getDetailCtx, stopGetDetail, consts.ProxyIpArr[i])
	}
	//4.SecKill
	var secKillWg sync.WaitGroup
	seckillCtx, stopSeckill := context.WithCancel(context.Background())
	for {
		select {
		case <-seckillCtx.Done():
			goto END
		case detail := <-dateDetailCh:
			for i := 0; i < len(consts.ProxyIpArr); i++ {
				secKillWg.Add(1)
				go func(detail model.DateDetail, ip string) {
					log.Printf("[%s]: 日期信息获取成功，尝试预约：%s %s-%s \n", ip, detail.Date, detail.StartTime, detail.EndTime)
					e.SecKill(detail, strconv.Itoa(productId), secKillWg,
						seckillCtx, stopGetDetail, stopSeckill, ip)
				}(detail, consts.ProxyIpArr[i])
			}
		default:
		}
	}
	secKillWg.Wait()
END:
	fmt.Printf("zmyy_seckill程序运行结束，共用时：%s, 按任意键退出...\n", time.Now().Sub(startTime))
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
