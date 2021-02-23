package zhimiaoyiyue

import (
	"fmt"
	"testing"
	"zmyy_seckill/model"
	"zmyy_seckill/util"
)

type CustomerList struct {
	E      []Customer `json:"list"`
	Status int        `json:"status"`
}
type Customer struct {
	Id   int    `json:"id"`
	Name string `json:"cname"`
}

func Test_transfer(t *testing.T) {
	s1 := `
			{
				"list":[
				{
					"id":1776,
					"cname":"西安市未央区张家堡社区卫生服务中心",
					"addr":"凤城七路八号",
					"SmallPic":"https://app.zhifeishengwu.com/img/none.png",
					"lat":34.33705700,
					"lng":108.94514500,
					"tel":"18991825131",
					"addr2":"凤城七路八号",
					"province":2375,
					"city":2376,
					"county":2377,
					"sort":1,
					"distance":907.52,
					"tags":[]
				}
			],
			"status":200
			}
`
	b := []byte(s1)
	v, err := util.Transfer2Model(b, model.CustomerList{})
	m := v.(model.CustomerList)
	if err != nil {
		t.Errorf("failed, err : %v", err)
	}
	fmt.Printf("%v", m)
}

func Test_GetCustomerList(t *testing.T) {
	e := ZMYYEngine{}
	e.Init()
	list, err := e.GetCustomerList()
	if err != nil {
		t.Errorf("err : %v", err)
	}
	fmt.Printf("%v", list)

}
