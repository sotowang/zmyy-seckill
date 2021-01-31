package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"zmyy_seckill/src/bili/consts"
	"zmyy_seckill/src/bili/util"
)

func Fetch(url string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	contents, err := ioutil.ReadAll(resp.Body)

	if err == nil && consts.SessionId == "" {
		sessionIdstring := resp.Header.Get("Set-Cookie")
		consts.SessionId = util.ParseSessionId(sessionIdstring)
		fmt.Printf("Sessionid string got  : %s \n", consts.SessionId)
	}
	return contents, nil
}
