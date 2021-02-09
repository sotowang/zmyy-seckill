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
	e.Run()
	fmt.Println("运行结束.....")
}
