package zhimiaoyiyue

import (
	"fmt"
	"strconv"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/model"
	"zmyy_seckill/util"
)

/**
获取验证码图片
*/
func (e *ZMYYEngine) GetVerifyPic() error {
	url := consts.GetCaptchaUrl
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["Cookie"] = e.Conf.Cookie
	zftsl, err := util.GetZFTSL()
	headers["zftsl"] = zftsl
	if err != nil {
		return err
	}
	if err != nil {
		fmt.Printf("GetVerifyPic().getZftsl() err :%v\n", err)
		return err
	}
	headers["zftsl"] = zftsl
	err = fetcher.FetchBigResp(url, headers)
	if err != nil {
		return err
	}
	return nil
}

/**
滑块验证码验证
*/
func (e *ZMYYEngine) CaptchaVerify() (*model.VerifyResultModel, error) {
	//1.获取验证码的base64编码
	err := e.GetVerifyPic()
	if err != nil {
		return nil, err
	}
	//2.base64编码转图片
	err = util.Base64ToPics()
	if err != nil {
		return nil, err
	}
	tigerPath, dragonPath, processPath := "../imgs/tiger.png", "../imgs/dragon.png", "../imgs/process.png"
	//3.图片验证码识别
	x, err := util.CallPythonScript(tigerPath, dragonPath, processPath)
	if err != nil {
		return nil, err
	}
	url := consts.CaptchaVerifyUrl + "&token=&x=" + x + "&y=5"

	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["Cookie"] = e.Conf.Cookie
	zftsl, err := util.GetZFTSL()
	headers["zftsl"] = zftsl
	bytes, err := fetcher.Fetch(url, headers)
	if err != nil {
		fmt.Printf("CaptchaVerify() err :%v\n", err)
		return nil, err
	}
	m := &model.VerifyResultModel{}
	err = util.Transfer2VerifyResultModel(bytes, m)
	if err != nil || m.Status != 200 || m.Guid == "" {
		fmt.Printf("CaptchaVerify() 验证码识别失败 :%v\n", err)
		return nil, err
	}
	return m, nil
}

func (e *ZMYYEngine) Save20(date string) {
	url := consts.SaveUrl + "&birthday=" + e.Conf.Birthday + "&tel=" + e.Conf.Tel + "&sex=" + strconv.Itoa(e.Conf.Sex) + "&cname=" + util.UrlEncode(e.Conf.Name) + "&doctype=1&idcard=" + e.Conf.IdCard + "&mxid=" + e.Conf.Mxid + "&date=" + date + "&pid=" + e.Conf.Product + "&Ftime=1&guid="
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["cookie"] = e.Conf.Cookie
	headers["Connection"] = "keep-alive"
	bytes, err := fetcher.Fetch(url, headers)
	if err != nil {
		fmt.Printf("Save20() err : %v \n", err)
	}
	fmt.Printf("%s", bytes)
	return
}
