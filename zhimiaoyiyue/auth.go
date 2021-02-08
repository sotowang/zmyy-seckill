package zhimiaoyiyue

import (
	"fmt"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/util"
)

func (e *ZMYYEngine) AuthAndSetSessionID() error {
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	zftsl, _ := util.GetZFTSL()
	headers["zftsl"] = zftsl
	contents, err := fetcher.Fetch(consts.AuthUrl, headers)
	if err != nil {
		fmt.Printf("AuthAndSetSessionID err : %v \n", err)
		return err
	}
	fmt.Printf("%s \n", contents)
	return nil
}
