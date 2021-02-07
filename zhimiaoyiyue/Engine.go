package zhimiaoyiyue

import (
	"fmt"
	"os"
	"strconv"
	"zmyy_seckill/config"
	"zmyy_seckill/model"
)

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
	subscribeDates, _ := e.GetCustSubscribeDateAll(customerId, productId, e.Conf.Month)

	fmt.Printf("可预约时间如下：将尝试从以下时间中预约！ \n")
	for i, v := range subscribeDates.Dates {
		fmt.Printf("时间%d : %v\n", i+1, v.Date)
		//e.SaveOrder(v.Date,strconv.Itoa(productId))
	}
	e.SaveOrder("2021-02-10", strconv.Itoa(productId))
	fmt.Printf("Press any key to exit...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
