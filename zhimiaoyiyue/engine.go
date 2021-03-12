package zhimiaoyiyue

import (
	"context"
	"fmt"
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

func (e *ZMYYEngine) Init() {
	var c config.RootConf
	conf, err2 := c.GetConf()
	if err2 != nil {
		panic(err2)
	}
	e.Conf = conf
}

func (e *ZMYYEngine) Run(customerId, productId int) {
	detailOk := make(chan *model.SubscribeDate, 100)
	ctx, cancel := context.WithCancel(context.Background())
	var subscribeDates *model.SubscribeDate
	//获取疫苗可预约的日期
	for {
		var err error
		subscribeDates, err = e.GetCustSubscribeDateAll(customerId, productId, e.Conf.Month)
		if err != nil || len(subscribeDates.Dates) == 0 {
			if len(subscribeDates.Dates) == 0 {
				fmt.Printf("目前可预约日期：%d个,尝试重新获取日期...\n", len(subscribeDates.Dates))
			}
		} else {
			detailOk <- subscribeDates
			break
		}
	}
	//获取可预约日期内疫苗信息
	for {
		select {
		case <-ctx.Done():
			goto END
		default:
			visited := false
			for _, v := range subscribeDates.Dates {
				if v.Enable == false {
					continue
				}
				wg2.Wait()
				if consts.Stop {
					goto END
				}
				fmt.Printf("尝试获取 %s 疫苗信息...\n", v.Date)
				m, err := e.GetCustSubscribeDateDetail(v.Date, productId, customerId)
				if err != nil || len(m.DateDetails) == 0 {
					fmt.Printf("未找到 %s 的可预约时间，尝试查找下一个时间...\n", v.Date)
					continue
				} else {
					func(dateDetails model.SubscribeDateDetail) {
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
			//如果没进入秒杀程序，说明已被抢完了
			if !visited {
				fmt.Printf("在所找到的可预约日期中，没有可预约的疫苗了（说明已经被抢完了）...\n")
				goto END
			}
		}
	}
	wg.Wait()
END:
	fmt.Printf("按任意键退出程序...\n")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
