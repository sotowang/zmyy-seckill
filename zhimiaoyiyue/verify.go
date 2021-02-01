package zhimiaoyiyue

import (
	"fmt"
	"strconv"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/model"
	"zmyy_seckill/util"
)

func (e *ZMYYEngine) GetVerifyPic() {
	url := consts.GetCaptchaUrl
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["cookie"] = "ASP.NET_SessionId=" + consts.SessionId
	bytes, err := fetcher.Fetch(url, headers)
	if err != nil {
		fmt.Printf("GetCustomerProduct() err : %v \n", err)
	}
	subsDates := model.SubscribeDate{}
	err = util.TransferToSubscribeDateModel(bytes, &subsDates)
	if err != nil {
		fmt.Printf("GetCustSubscribeDateAll() err: %v\n ", err)

	}
	return

}
func (e *ZMYYEngine) CaptchaVerifyUrl() {
	url := consts.CaptchaVerifyUrl
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["cookie"] = "ASP.NET_SessionId=" + consts.SessionId
	bytes, err := fetcher.Fetch(url, headers)
	if err != nil {
		fmt.Printf("GetCustomerProduct() err : %v \n", err)
	}
	subsDates := model.SubscribeDate{}
	err = util.TransferToSubscribeDateModel(bytes, &subsDates)
	if err != nil {
		fmt.Printf("GetCustSubscribeDateAll() err: %v\n ", err)

	}
	return
}

func (e *ZMYYEngine) Save20(date string) {
	url := consts.SaveUrl + "&birthday=" + e.Conf.Birthday + "&tel=" + e.Conf.Tel + "&sex=" + strconv.Itoa(e.Conf.Sex) + "&cname=" + util.UrlEncode(e.Conf.Name) + "&doctype=1&idcard=" + e.Conf.IdCard + "&mxid=" + e.Conf.Mxid + "&date=" + date + "&pid=" + e.Conf.Product + "&Ftime=1&guid=" + e.Conf.Guid
	url = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=Save20&birthday=1995-03-01&tel=17630379886&sex=1&cname=%E7%8E%8B%E6%9D%BE%E6%B6%9B&doctype=1&idcard=412723199503012532&mxid=AAAAACJZAAAcYjQB&date=2021-02-04&pid=54&Ftime=1&guid=4f300e61-4eae-4bee-9315-8f5b43c3bff9"
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["cookie"] = "ASP.NET_SessionId=" + consts.SessionId
	headers["Connection"] = "keep-alive"
	bytes, err := fetcher.Fetch(url, headers)
	if err != nil {
		fmt.Printf("Save20() err : %v \n", err)
	}
	fmt.Printf("%s", bytes)
	return
}
