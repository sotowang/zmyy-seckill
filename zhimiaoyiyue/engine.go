package zhimiaoyiyue

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
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

var wg sync.WaitGroup
var wg2 sync.WaitGroup
var mutex sync.Mutex
var flag bool
var stop bool

func (e *ZMYYEngine) Init() {
	var c config.RootConf
	conf, err2 := c.GetConf()
	if err2 != nil {
		panic(err2)
	}
	e.Conf = conf
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
				fmt.Printf("未获取到可预约日期,尝试重新获取...\n")
			} else {
				subscribeDateChan <- subscribeDates
				stop()
			}
		}

	}
}

/**
获取筛选后的日期的可预约时间段
*/
func getDateDetail(availableDates []model.Dates, productId int, customerId int, e *ZMYYEngine, ctx context.Context, stop context.CancelFunc, ip ...string) {
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
			mutex.Lock()
			index := rand.Intn(len(availableDates))
			v := availableDates[index]
			mutex.Unlock()
			fmt.Printf("尝试获取 %s 疫苗信息...\n", v.Date)
			dateDetails := e.GetCustSubscribeDateDetail(v.Date, productId, customerId)
			if dateDetails == nil || len(dateDetails.DateDetails) == 0 {
				fmt.Printf("未找到 %s 的可预约时间，尝试查找下一个时间...\n", v.Date)
			} else {

			}
		}

	}
}

func (e *ZMYYEngine) Run(customerId, productId int) {
	//var subscribeDates *model.SubscribeDate
	availableDates := make([]model.Dates, 0)
	var subscribeDateChan = make(chan *model.SubscribeDate)

	//1. 使用多个ip多线程获取疫苗可预约的日期
	getDatesCtx, stopGetDates := context.WithCancel(context.Background())
	for i := 0; i < len(consts.ProxyIpArr); i++ {
		go getAvailableDates(customerId, productId, e.Conf.Month, subscribeDateChan, e, getDatesCtx, stopGetDates, consts.ProxyIpArr[i])
	}
	log.Println("正在等待获取日期...")
	subscribeDates := <-subscribeDateChan
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
	for len(availableDates) > 0 {
		k := 0
		for i := 0; i < len(availableDates); i++ {
			if stop {
				goto END
			}
			for flag {
			}
			v := availableDates[i]
			fmt.Printf("尝试获取 %s 疫苗信息...\n", v.Date)
			dateDetails := e.GetCustSubscribeDateDetail(v.Date, productId, customerId)
			if dateDetails == nil || len(dateDetails.DateDetails) == 0 {
				fmt.Printf("未找到 %s 的可预约时间，尝试查找下一个时间...\n", v.Date)
				if dateDetails == nil {
					//如果未拿到availableDates[i]的数据，则下一轮循环继续查找
					availableDates[k] = v
					k++
				}
				continue
			} else {
				for _, v := range dateDetails.DateDetails {
					if v.Qty <= 0 {
						continue
					}
					v.Date = dateDetails.Date
					wg.Add(1)
					go func(detail model.DateDetail) {
						fmt.Printf("日期信息获取成功，尝试预约：%s %s-%s \n", detail.Date, detail.StartTime, detail.EndTime)
						e.SecKill(detail, strconv.Itoa(productId))
					}(v)
				}
			}
		}
		availableDates = availableDates[:k]
	}
	wg.Wait()
END:
	fmt.Printf("zmyy_seckill程序运行结束，按任意键退出...\n")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
