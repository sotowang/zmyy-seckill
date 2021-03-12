package consts

import "zmyy_seckill/limit"

var RequestLimitRate limit.LimitRate
var Stop bool

const (
	UserAgent  = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat"
	Refer      = "https://servicewechat.com/wx2c7f0f3c30d99445/72/page-frame.html"
	Connection = "keep-alive"
	//Host       = "https://106.13.187.42"
	Host = "https://cloud.cn2030.com"
	//Host       = "https://183.230.139.228"
	//Host       = "https://yun.cn2030.com"
	//某地区医院列表URL
	CustomerListUrl = Host + "/sc/wx/HandlerSubscribe.ashx?act=CustomerList"
	//授权URL
	AuthUrl = Host + "/sc/wx/HandlerSubscribe.ashx?act=auth&code=061H55000QOs8L1yHN100Ba0N43H550I"
	//某医院内HPV疫苗情况URL
	CustomerProductURL = Host + "/sc/wx/HandlerSubscribe.ashx?act=CustomerProduct"
	//预约用户信息
	UserInfoURL          = Host + "/sc/wx/HandlerSubscribe.ashx?act=User"
	CustSubscribeDateUrl = Host + "/sc/wx/HandlerSubscribe.ashx?act=GetCustSubscribeDateAll"
	CaptchaVerifyUrl     = Host + "/sc/wx/HandlerSubscribe.ashx?act=CaptchaVerify"
	GetCaptchaUrl        = Host + "/sc/wx/HandlerSubscribe.ashx?act=GetCaptcha"
	SaveUrl              = Host + "/sc/wx/HandlerSubscribe.ashx?act=Save20"
	//获取订单状态
	OrderStatusUrl             = Host + "/sc/wx/HandlerSubscribe.ashx?act=GetOrderStatus"
	CustSubscribeDateDetailUrl = Host + "/sc/wx/HandlerSubscribe.ashx?act=GetCustSubscribeDateDetail"
)
