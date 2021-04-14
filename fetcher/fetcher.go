package fetcher

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
	"zmyy_seckill/consts"
	"zmyy_seckill/utils"
)

//使用限流器请求url
func FetchWithRatelimter(zhimiaoUrl string, headers map[string]string, v ...string) ([]byte, error) {
	consts.RequestLimitRate.Limit()
	return Fetch(zhimiaoUrl, headers, v...)
}

//常规请求url
func Fetch(zhimiaoUrl string, headers map[string]string, v ...string) ([]byte, error) {
	//var proxy *url.URL
	var proxy func(*http.Request) (*url.URL, error)
	proxyIp := ""
	//目前v参数中为代理ip，若v长度>0，则需要使用代理
	for _, ip := range v {
		proxyIp = ip
	}
	if proxyIp != "" {
		proxyUrl, _ := url.Parse(proxyIp)
		proxy = http.ProxyURL(proxyUrl)
	}
	log.Printf("%s 正在发起请求.... url: %s\n", proxyIp, zhimiaoUrl)
	tr := &http.Transport{
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		Proxy:                 proxy,
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * 3,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}
	req, err := http.NewRequest("GET", zhimiaoUrl, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("%s fetch err : %v \n", proxyIp, err)
		//utils.SetRandomIP()
		return nil, err
	}
	//切换IP
	defer resp.Body.Close()
	//如果有重定向错误，则重定向
	if resp.Request.Response != nil && resp.Request.Response.StatusCode == http.StatusFound {
		zhimiaoUrl = consts.Host + resp.Request.Response.Header.Get("Location")
		return Fetch(zhimiaoUrl, headers, v...)
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("%s fetch err : %v \n", proxyIp, resp.StatusCode)
		//切换IP
		//utils.SetRandomIP()
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%s fetch err : %v \n", proxyIp, err)
		//切换IP
		//utils.SetRandomIP()
		return nil, err
	}
	return contents, nil
}
func FetchCaptcha(captchaUrl string, headers map[string]string, prefix string, v ...string) error {
	consts.RequestLimitRate.Limit()
	var proxy *url.URL
	proxyIp := ""
	if len(v) > 0 {
		proxyIp = v[0]
		proxy, _ = url.Parse(proxyIp)
	}
	log.Printf("%s 正在发起请求.... captchaUrl: %s\n", proxyIp, captchaUrl)
	tr := &http.Transport{
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * 3,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}
	req, err := http.NewRequest("GET", captchaUrl, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("%s fetch err : %v \n", proxyIp, err)
		//切换IP
		utils.SetRandomIP()
		return err
	}
	//如果有重定向错误，则重定向
	if resp.Request.Response != nil && resp.Request.Response.StatusCode == http.StatusFound {
		captchaUrl = consts.Host + resp.Request.Response.Header.Get("Location")
		return FetchCaptcha(captchaUrl, headers, prefix)
	}
	b, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil || resp.StatusCode != http.StatusOK || b < 100 {
		fmt.Printf("获取验证码图片失败，请求可能被禁止！code： %d\n", resp.StatusCode)
		log.Printf("%s fetch err : %v \n", proxyIp, err)
		utils.SetRandomIP()
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
