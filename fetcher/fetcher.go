package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"zmyy_seckill/util"
)

var rateLimiter = time.Tick(1000 * time.Millisecond)

func Fetch(url string, headers map[string]string) ([]byte, error) {
	<-rateLimiter
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
	if err != nil {
		return nil, err
	}
	return contents, nil
}
func FetchBigResp(url string, headers map[string]string, prefix string) error {
	<-rateLimiter
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Printf("wrong status code : %d\n", resp.StatusCode)
		return err
	}
	path := util.GetCurrentPath()
	path = path + "/imgs/" + prefix
	f, _ := os.Create(path)
	defer f.Close()

	buf := make([]byte, 1024*1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		f.Write(buf[:n])
	}
	return nil
}
