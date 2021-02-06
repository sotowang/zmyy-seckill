package fetcher

import (
	"fmt"
	"net/http"
	"zmyy_seckill/consts"
	"zmyy_seckill/util"
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
	//contents, err := ioutil.ReadAll(resp.Body)
	buf := make([]byte, 4096)
	result := ""
	for {
		n, err := resp.Body.Read(buf)
		if err != nil || n == 0 {
			break
		}
		result += string(buf[:n])
	}
	//contents, err := ioutil.ReadAll(resp.Body)

	if err == nil && consts.SessionId == "" {
		sessionIdstring := resp.Header.Get("Set-Cookie")
		consts.SessionId = util.ParseSessionId(sessionIdstring)
		fmt.Printf("Sessionid string got  : %s \n", consts.SessionId)
	}
	return nil, nil
}
