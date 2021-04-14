package main

import (
	"fmt"
	"time"
	"zmyy_seckill/consts"
	"zmyy_seckill/ip"
	"zmyy_seckill/zhimiaoyiyue"
)

var stop bool

func main() {
	e := zhimiaoyiyue.ZMYYEngine{}
	e.Init()
	//获取可用代理ip,下行代码开启时则启用ip代理,默认使用本机的ip
	consts.ProxyIpArr = ip.ReadIpFile()
	//utils.SetRandomIP()
	customerId, productId := -1, -1
	//设置抢购请求速率，2s/次，下行代码开启时则开始限流
	consts.RequestLimitRate.SetRate(1, 20)
	for customerId == -1 || productId == -1 {
		if customerId == -1 {
			//获取指定地区接种地点的customerId
			customerId, _ = e.GetCustomerList()
		}
		if productId == -1 {
			//获取指定接种地点的productId
			productId, _ = e.GetCustomerProduct(customerId)
		}
	}
	loc, _ := time.LoadLocation("Local")
	timeLayout := "2006-01-02 15:04:05"
	subsTime, _ := time.ParseInLocation(timeLayout, e.Conf.SubscribeTime, loc)
	now := time.Now()
	timer := time.NewTimer(subsTime.Sub(now))
	fmt.Printf("倒计时中，将在 %v 时运行程序...\n", subsTime)
	<-timer.C
	fmt.Printf("开始运行zmyy_seckill....\n")
	e.Run(customerId, productId)
}
