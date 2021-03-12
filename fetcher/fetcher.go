package fetcher

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"zmyy_seckill/consts"
	"zmyy_seckill/util"
)

func Fetch(url string, headers map[string]string) ([]byte, error) {
	consts.RequestLimitRate.Limit()
	fmt.Printf("正在发起请求.... url: %s\n", url)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("GET", url, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//如果有重定向错误，则重定向
	if resp.Request.Response != nil && resp.Request.Response.StatusCode == http.StatusFound {
		url = consts.Host + resp.Request.Response.Header.Get("Location")
		//fmt.Printf("出现302错误，尝试重定向网址...\n")
		return Fetch(url, headers)
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
	consts.RequestLimitRate.Limit()
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	//如果有重定向错误，则重定向
	if resp.Request.Response != nil && resp.Request.Response.StatusCode == http.StatusFound {
		url = consts.Host + resp.Request.Response.Header.Get("Location")
		fmt.Printf("出现302错误，尝试重定向网址...\n")
		return FetchBigResp(url, headers, prefix)
	}
	defer resp.Body.Close()
	b, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil || resp.StatusCode != http.StatusOK || b < 100 {
		fmt.Printf("获取验证码图片失败，请求可能被禁止！code： %d\n", resp.StatusCode)
		return errors.New("获取验证码图片失败，请求可能被禁止！")
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
