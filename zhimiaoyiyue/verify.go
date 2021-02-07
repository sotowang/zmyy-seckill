package zhimiaoyiyue

import (
	"fmt"
	"time"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/model"
	"zmyy_seckill/util"
)

/**
获取验证码Base64编码
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
	time.Sleep(1 * time.Second)
	if err != nil {
		return nil, err
	}
	//2.base64编码转图片
	err = util.Base64ToPics()
	if err != nil {
		return nil, err
	}
	path := util.GetCurrentPath()
	tigerPath, dragonPath, processPath := path+"/imgs/tiger.png", path+"/imgs/dragon.png", path+"/imgs/process.png"
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
		fmt.Printf("CaptchaVerify() 验证码识别失败 :err:%v, %s\n", err, bytes)
		return nil, err
	}
	return m, nil
}
