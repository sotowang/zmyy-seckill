package zhimiaoyiyue

import (
	"fmt"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
)

func AuthAndSetSessionID() error {
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	contents, err := fetcher.Fetch(consts.AuthUrl, headers)
	if err != nil {
		fmt.Printf("err : %v \n", err)
		return err
	}
	fmt.Printf("%s \n", contents)
	return nil
}
