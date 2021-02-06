package zhimiaoyiyue

import (
	"fmt"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/util"
)

func (e *ZMYYEngine) getUserInfo() {
	UserInfoURL := consts.UserInfoURL
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	//setSessionId
	AuthAndSetSessionID()
	headers["cookie"] = "ASP.NET_SessionId=" + consts.SessionId
	zftsl, err2 := util.CallJsScript()
	if err2 != nil {
		return
	}
	headers["zftsl"] = zftsl
	contents, err := fetcher.Fetch(UserInfoURL, headers)
	if err != nil {
		return
	}
	fmt.Printf("%s", contents)
}
