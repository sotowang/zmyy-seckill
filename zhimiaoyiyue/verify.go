package zhimiaoyiyue

import (
	"fmt"
	"strings"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/model"
	"zmyy_seckill/util"
)

/**
获取验证码图片
*/
func (e *ZMYYEngine) GetVerifyPic(dateDetail model.DateDetail) (path string, err error) {
	url := consts.GetCaptchaUrl
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["Cookie"] = e.Conf.Cookie
	headers["Connection"] = consts.Connection
	zftsl, _ := util.GetZFTSL()
	headers["zftsl"] = zftsl
	prefix := dateDetail.Date + "-" + strings.Replace(dateDetail.StartTime, ":", "_", -1) + "-" + strings.Replace(dateDetail.EndTime, ":", "_", -1)
	err = fetcher.FetchBigResp(url, headers, prefix)
	if err != nil {
		return "", err
	}
	path = util.GetCurrentPath() + "/imgs/"
	//Base64转图片
	err = util.Base64ToPics(path + prefix)
	if err != nil {
		return "", err
	}
	return path + prefix, nil
}

/**
滑块验证码验证
*/
func (e *ZMYYEngine) CaptchaVerify(prefix string) (guid string, err error) {
	tigerPath, dragonPath, processPath := prefix+"-tiger.png", prefix+"-dragon.png", prefix+"-process.png"
	//2.图片验证码识别
	x, err := util.CallPythonScript(tigerPath, dragonPath, processPath)

	if err != nil {
		return "", err
	}
	url := consts.CaptchaVerifyUrl + "&token=&x=" + x + "&y=5"

	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["Cookie"] = e.Conf.Cookie
	headers["Connection"] = consts.Connection
	zftsl, err := util.GetZFTSL()
	headers["zftsl"] = zftsl
	bytes, err := fetcher.Fetch(url, headers)
	if err != nil {
		return "", err
	}
	m := &model.VerifyResultModel{}
	err = util.Transfer2VerifyResultModel(bytes, m)
	defer util.DeleteFile(tigerPath, dragonPath, processPath, prefix)
	if err != nil || m.Status != 200 || m.Guid == "" {
		fmt.Printf("CaptchaVerify() 验证码%s验证失败，guid=%s ; err:%v; %s\n", prefix, m.Guid, err, bytes)
		return "", err
	}

	return m.Guid, nil
}
