package zhimiaoyiyue

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
	"zmyy_seckill/config"
	"zmyy_seckill/model"
)

var wg sync.WaitGroup

type ZMYYEngine struct {
	Conf config.CustomerConf
}

type SecKill interface {
	GetCustomerList() (*model.CustomerList, error)
}

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
	//获取可预约的时间
	dateOk := make(chan bool, 1)
	dateOk <- true
	dateChan := make(chan model.DateDetail, 100)
	for {
		select {
		case <-dateOk:
			subscribeDates, err := e.GetCustSubscribeDateAll(customerId, productId, e.Conf.Month)
			if err != nil || len(subscribeDates.Dates) == 0 {
				fmt.Printf("目前可预约日期：%d个\n", len(subscribeDates.Dates))
				dateOk <- true
			} else if len(subscribeDates.Dates) > 0 {
				//获取这可预约日期的具体信息，包括疫苗数量等
				for _, v := range subscribeDates.Dates {
					//获取该日期的具体信息
					m, err := e.GetCustSubscribeDateDetail(v.Date, productId, customerId)
					if err != nil {
						dateOk <- true
					} else {
						for _, v := range m.DateDetails {
							if v.Qty > 0 {
								v.Date = m.Date
								dateChan <- v
							}
						}
						if len(dateChan) > 0 {
							goto START
						} else {
							dateOk <- true
						}
					}
				}
			}
		}
	}
START:
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Printf("可预约时间如下：将尝试从以下时间中预约 \n")
	for {
		select {
		case dateDetail := <-dateChan:
			wg.Add(1)
			//index := strconv.FormatInt(time.Now().Unix(), 10)
			time.Sleep(time.Second)
			go e.Seckill(ctx, cancel, dateDetail, strconv.Itoa(productId), "")
		}
	}
	wg.Wait()

	fmt.Printf("订单抢购成功！Press any key to exit...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
