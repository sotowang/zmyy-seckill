package zhimiaoyiyue

import (
	"context"
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

func (e *ZMYYEngine) Init() {
	var c config.RootConf
	conf, err2 := c.GetConf()
	if err2 != nil {
		panic(err2)
	}
	e.Conf = conf
}

func (e *ZMYYEngine) Run(customerId, productId int) {

	dateOk := make(chan bool, 1)
	detailOk := make(chan *model.SubscribeDate, 100)
	//dateChan := make(chan model.DateDetail, 100)
	ctx, cancel := context.WithCancel(context.Background())
	dateOk <- true
	for {
		select {
		//获取可预约的日期
		case <-dateOk:
			subscribeDates, err := e.GetCustSubscribeDateAll(customerId, productId, e.Conf.Month)
			if err != nil || len(subscribeDates.Dates) == 0 {
				if err == nil {
					fmt.Printf("目前可预约日期：%d个,尝试重新获取日期...\n", len(subscribeDates.Dates))
				}
				dateOk <- true
			} else {
				detailOk <- subscribeDates
			}
		case dates := <-detailOk:
			//temDates := dates
			visited := false
			//获取这可预约日期的具体信息，包括疫苗数量等
			for _, v := range dates.Dates {
				if v.Enable == false {
					continue
				}
				wg2.Wait()
				//获取该日期的具体信息
				fmt.Printf("尝试获取 %s 疫苗信息...\n", v.Date)
				m, err := e.GetCustSubscribeDateDetail(v.Date, productId, customerId)
				if err != nil || len(m.DateDetails) == 0 {
					fmt.Printf("未找到 %s 的可预约时间，尝试查找下一个时间...\n", v.Date)
					//detailOk <- temDates
					continue
				} else {
					go func(dateDetails model.SubscribeDateDetail) {
						for _, v := range dateDetails.DateDetails {
							if v.Qty > 0 {
								visited = true
								v.Date = dateDetails.Date
								wg.Add(1)
								go func(detail model.DateDetail) {
									fmt.Printf("日期信息获取成功，尝试预约：%s %s-%s \n", detail.Date, detail.StartTime, detail.EndTime)
									e.SecKill(ctx, cancel, detail, strconv.Itoa(productId))
								}(v)
							}
						}
					}(*m)
				}
			}
			if !visited {
				fmt.Printf("未找到可预约时间，尝试重新获取可预约日期...\n")
				//detailOk <- temDates
				dateOk <- true
			}
		case <-ctx.Done():
			goto END
		default:
		}
	}
	wg.Wait()
END:
	fmt.Printf("订单抢购成功！Press any key to exit...\n")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
