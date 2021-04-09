package fetcher

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
	"zmyy_seckill/consts"
	"zmyy_seckill/ip"
	"zmyy_seckill/utils"
)

func Fetch(zhimiaoUrl string, headers map[string]string) ([]byte, error) {
	consts.RequestLimitRate.Limit()
	var proxy *url.URL
	if consts.ProxyIp != "" {
		proxy, _ = url.Parse(consts.ProxyIp)
	}
	fmt.Printf("ip: %s 正在发起请求.... url: %s\n", consts.ProxyIp, zhimiaoUrl)
	tr := &http.Transport{
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(3),
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   3 * time.Second,
	}
	req, err := http.NewRequest("GET", zhimiaoUrl, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		ip.SetRandomIP()
		return nil, err
	}
	//切换IP
	defer resp.Body.Close()
	//如果有重定向错误，则重定向
	if resp.Request.Response != nil && resp.Request.Response.StatusCode == http.StatusFound {
		zhimiaoUrl = consts.Host + resp.Request.Response.Header.Get("Location")
		return Fetch(zhimiaoUrl, headers)
	}
	if resp.StatusCode != http.StatusOK {
		//切换IP
		ip.SetRandomIP()
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//切换IP
		ip.SetRandomIP()
		return nil, err
	}
	return contents, nil
}
func FetchBigResp(captchaUrl string, headers map[string]string, prefix string) error {
	consts.RequestLimitRate.Limit()
	fmt.Printf("ip: %s 正在发起请求.... captchaUrl: %s\n", consts.ProxyIp, captchaUrl)
	var proxy *url.URL
	if consts.ProxyIp != "" {
		proxy, _ = url.Parse(consts.ProxyIp)
	}
	tr := &http.Transport{
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(3),
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   3 * time.Second,
	}
	req, err := http.NewRequest("GET", captchaUrl, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		//切换IP
		ip.SetRandomIP()
		return err
	}
	//如果有重定向错误，则重定向
	if resp.Request.Response != nil && resp.Request.Response.StatusCode == http.StatusFound {
		captchaUrl = consts.Host + resp.Request.Response.Header.Get("Location")
		return FetchBigResp(captchaUrl, headers, prefix)
	}
	b, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil || resp.StatusCode != http.StatusOK || b < 100 {
		fmt.Printf("获取验证码图片失败，请求可能被禁止！code： %d\n", resp.StatusCode)
		ip.SetRandomIP()
		return errors.New("获取验证码图片失败，请求可能被禁止！")
	}
	defer resp.Body.Close()
	path := utils.GetCurrentPath()
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
