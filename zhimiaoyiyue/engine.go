package zhimiaoyiyue

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"zmyy_seckill/config"
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

func (e *ZMYYEngine) Run(customerId, productId int) {
	var subscribeDates *model.SubscribeDate
	availableDates := make([]model.Dates, 0)
	//获取疫苗可预约的日期
	for len(availableDates) == 0 {
		var err error
		subscribeDates, err = e.GetCustSubscribeDateAll(customerId, productId, e.Conf.Month)
		if err != nil || len(subscribeDates.Dates) == 0 {
			fmt.Printf("未获取到可预约日期,尝试重新获取...\n")
		} else {
			for _, v := range subscribeDates.Dates {
				if v.Enable == false {
					continue
				}
				availableDates = append(availableDates, v)
			}
			break
		}
	}
	//获取可预约日期内疫苗信息并预约
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
			dateDetails, err := e.GetCustSubscribeDateDetail(v.Date, productId, customerId)
			if err != nil || len(dateDetails.DateDetails) == 0 {
				fmt.Printf("未找到 %s 的可预约时间，尝试查找下一个时间...\n", v.Date)
				if err != nil {
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
	fmt.Printf("任务结束，按任意键退出程序...\n")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
