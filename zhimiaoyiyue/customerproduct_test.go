package zhimiaoyiyue

import (
	"fmt"
	"testing"
)

func TestCustomerProduct(t *testing.T) {
	//bytes := `
	//	{
	//		"status":200,
	//		"startDate":"2021-02-03 09:00:00",
	//		"endDate":"2021-02-03 20:00:00",
	//		"memo":"",
	//		"tel":"18991825131",
	//		"addr":"凤城七路八号",
	//		"cname":"西安市未央区张家堡社区卫生服务中心",
	//		"lat":34.33705700,
	//		"lng":108.94514500,
	//		"distance":11749.810546875,
	//		"payment":{
	//			"alipay":"true",
	//			"WechatPay":"true",
	//			"UnionPay":"true",
	//			"cashier":"true"
	//		},
	//		"BigPic":"https://app.zhifeishengwu.com/img/none.png",
	//		"IdcardLimit":true,
	//		"notice":"",
	//		"list":[
	//			{
	//				"id":1,
	//				"text":"九价人乳头瘤病毒疫苗",
	//				"price":1338.00,
	//				"descript":"九价简介：接种九价宫颈癌疫苗可以预防人乳头瘤病毒6、11、16、18、31、33、45、52、58型感染引起的尖锐湿疣、癌前病变、原位腺癌和宫颈癌，是目前覆盖HPV型别最多的HPV疫苗。全程三针， 适用于16-26岁女性。",
	//				"warn":"适用于16-26岁女性。",
	//				"tags":["进口","默沙东","16-26","3针次"],
	//				"questionnaireId":0,
	//				"remarks":"",
	//				"NumbersVaccine":[
	//					{
	//						"cname":"第一针",
	//						"value":1},
	//					{
	//						"cname":"第二针",
	//						"value":2},
	//					{
	//						"cname":"第三针",
	//						"value":3
	//					}
	//				],
	//				"date":"02-03 09:00 至 02-03 20:00",
	//				"BtnLable":"暂未开始",
	//				"enable":false
	//			}
	//		]
	//	}
	//`
	//rootSource := model.RootSource{}
	//err := util.TransferToCustomerProductListModel([]byte(bytes), &rootSource)
	e := ZMYYEngine{}
	product, err := e.GetCustomerProduct(1776)
	if err != nil {
		t.Errorf("err : %v", err)
	}
	fmt.Printf("%v", product)
}
