package zhimiaoyiyue

import (
	"fmt"
	"strconv"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/model"
	"zmyy_seckill/util"
)

func (e *ZMYYEngine) GetVerifyPic(zftsl string) (*model.VerifyPicModel, error) {
	url := consts.GetCaptchaUrl
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["cookie"] = "ASP.NET_SessionId=" + consts.SessionId
	headers["zftsl"] = zftsl
	bytes, err := fetcher.Fetch(url, headers)
	if err != nil {
		fmt.Printf("GetCustomerProduct() err : %v \n", err)
		return nil, err
	}
	pics := model.VerifyPicModel{}
	err = util.TransferToVerifyModel(bytes, &pics)
	if err != nil {
		fmt.Printf("GetVerifyPic() err: %v\n ", err)
		return nil, err
	}
	return &pics, nil

}
func (e *ZMYYEngine) CaptchaVerify() {
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
