package zhimiaoyiyue

import (
	"fmt"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/utils"
)

func (e *ZMYYEngine) getUserInfo() {
	UserInfoURL := consts.UserInfoURL
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["Cookie"] = e.Conf.Cookie
	zftsl, err2 := utils.GetZFTSL()
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
