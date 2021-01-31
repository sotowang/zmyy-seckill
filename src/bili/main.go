package main

import (
	"fmt"
	"zmyy_seckill/src/bili/zhimiaoyiyue"
)

func main() {
	e := zhimiaoyiyue.ZMYYEngine{}
	e.Init()
	e.Run()
	fmt.Println("运行结束.....")
}
