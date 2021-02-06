package zhimiaoyiyue

import (
	"fmt"
	"zmyy_seckill/config"
	"zmyy_seckill/model"
)

type ZMYYEngine struct {
	Conf    config.CustomerConf
	Seckill SecKill
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
	//授权
	err := AuthAndSetSessionID()
	if err != nil {
		fmt.Printf("授权AuthAndSetSessionID（）出现未知错误！ err : %v \n", err)
	}
	//获取指定地区接种地点列表
	customerList, err := e.GetCustomerList()
	customers := customerList.Customers
	fmt.Printf("指定地方下，共找到 %d 个接种地点 ：\n", len(customers))
	for i, v := range customers {
		fmt.Printf("第%d个:  %v \n", i, v)
	}

	//获取ID这1776的接种地的信息
	products, err := e.GetCustomerProduct(e.Conf.CustomerId)
	fmt.Printf("以下为 --%s-- 信息: \n", products.Cname)
	fmt.Printf("开始预约日期：%s ; 结束预约日期 %s \n", products.StartDate, products.EndDate)
	fmt.Printf("疫苗信息：\n")
	cps := products.CustomerProducts
	for i, v := range cps {
		fmt.Printf("疫苗 %d ：id = %d, name = %s, price=%.2f 元\n", i+1, v.Id, v.Text, v.Price)
	}

}
