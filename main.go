package main

import (
	"fmt"
	"zmyy_seckill/consts"
	"zmyy_seckill/zhimiaoyiyue"
)

func main() {
	consts.RequestLimitRate.SetRate(1)
	e := zhimiaoyiyue.ZMYYEngine{}
	e.Init()
	customerId := -1
	productId := -1
	for customerId == -1 || productId == -1 {
		if customerId == -1 {
			//获取指定地区接种地点的cutomerId
			customerId, _ = e.GetCustomerList()
		}
		if productId == -1 {
			//获取指定接种地点的productId
			productId, _ = e.GetCustomerProduct(customerId)
		}
	}

	e.Run(customerId, productId)
	fmt.Println("运行结束.....")
}
