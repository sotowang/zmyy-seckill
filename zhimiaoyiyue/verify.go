package zhimiaoyiyue

import (
	"fmt"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/model"
	"zmyy_seckill/util"
)

/**
获取验证码图片
*/
func (e *ZMYYEngine) GetVerifyPic(date, productId string, index string) error {
	url := consts.GetCaptchaUrl
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["Cookie"] = e.Conf.Cookie
	zftsl, _ := util.GetZFTSL()
	headers["zftsl"] = zftsl
	prefix := productId + "-" + date + "-" + index
	err := fetcher.FetchBigResp(url, headers, prefix)
	if err != nil {
		fmt.Printf("获取验证码Base64编码失败：err : %v\n", err)
		return err
	}
	//Base64转图片
	err = util.Base64ToPics(prefix)
	if err != nil {
		return err
	}
	return nil
}

/**
滑块验证码验证
*/
func (e *ZMYYEngine) CaptchaVerify(date, productId string, index string) (guid string, err error) {
	path := util.GetCurrentPath()
	prefix := productId + "-" + date + "-" + index
	tigerPath, dragonPath, processPath := path+"/imgs/"+prefix+"tiger.png", path+"/imgs/"+prefix+"dragon.png", path+"/imgs/"+prefix+"process.png"
	//2.图片验证码识别
	x, err := util.CallPythonScript(tigerPath, dragonPath, processPath)
	util.DeleteFile(tigerPath, dragonPath, processPath, path+"/imgs/"+prefix)
	if err != nil {
		return "", err
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
		return "", err
	}
	m := &model.VerifyResultModel{}
	err = util.Transfer2VerifyResultModel(bytes, m)
	if err != nil || m.Status != 200 || m.Guid == "" {
		fmt.Printf("CaptchaVerify() 验证码%s-%s-%s验证失败，guid=%s ; err:%v; %s\n", productId, date, index, m.Guid, err, bytes)
		return "", err
	}
	return m.Guid, nil
}
