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

func (e *ZMYYEngine) Init() {
	var c config.RootConf
	conf, err2 := c.GetConf()
	if err2 != nil {
		panic(err2)
	}
	e.Conf = conf
}

func (e *ZMYYEngine) Run() {
	//获取指定地区接种地点的cutomerId
	customerId, _ := e.GetCustomerList()
	//获取指定接种地点的productId
	productId, _ := e.GetCustomerProduct(customerId)
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
				fmt.Printf("目前可预约日期：%d个,尝试重新获取日期...\n", len(subscribeDates.Dates))
				dateOk <- true
			} else {
				detailOk <- subscribeDates
			}
		case dates := <-detailOk:
			temDates := dates
			//获取这可预约日期的具体信息，包括疫苗数量等
			for _, v := range dates.Dates {
				//获取该日期的具体信息
				fmt.Printf("尝试获取%s 疫苗信息...\n", v.Date)
				m, err := e.GetCustSubscribeDateDetail(v.Date, productId, customerId)
				if err != nil {
					detailOk <- temDates
				} else {
					go func(dateDetails model.SubscribeDateDetail) {
						for _, v := range dateDetails.DateDetails {
							if v.Qty > 0 {
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
