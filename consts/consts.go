package consts

var (
	SessionId string
)

const (
	UserAgent = "Mozilla/5.0 (Linux; Android 7.1.2; TAS-AN00 Build/N2G47H; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.70 Mobile Safari/537.36 MicroMessenger/7.0.12.1620(0x27000C34) Process/appbrand0 NetType/WIFI Language/zh_CN ABI/arm64"
	Refer     = "https://servicewechat.com/wx2c7f0f3c30d99445/72/page-frame.html"
	//某地区医院列表URL
	CustomerListUrl = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=CustomerList"
	//授权URL
	AuthUrl = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=auth&code=061H55000QOs8L1yHN100Ba0N43H550I"
	//某医院内HPV疫苗情况URL
	CustomerProductURL = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=CustomerProduct"
	//预约用户信息
	UserInfoURL          = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=User"
	CustSubscribeDateUrl = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=GetCustSubscribeDateAll"
	CaptchaVerifyUrl     = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=CaptchaVerify"
	GetCaptchaUrl        = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=GetCaptcha"
	SaveUrl              = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=Save20"
)
