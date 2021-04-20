package ip

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
	"zmyy_seckill/consts"
	"zmyy_seckill/utils"
)

var ipCh = make(chan string, 100)

/**
判断代理ip的有效性
*/
func proxyTest(ip string, wg *sync.WaitGroup) {
	defer wg.Done()
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	//testUrl := "https://cloud.cn2030.com"
	testUrl := "https://www.baidu.com"
	// 解析代理地址
	proxy, err := url.Parse(ip)
	//设置网络传输
	tr := &http.Transport{
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * 3,
	}
	// 创建连接客户端
	client := &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}
	req, err := http.NewRequest("GET", testUrl, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return
	}
	defer resp.Body.Close()
	fmt.Printf("有效ip : %v \n", ip)
	ipCh <- ip
}
func ReadIpFile() (ipArr []string) {
	log.Printf("========================正在验证IP=============================\n")
	//defer close(ipCh)
	path := utils.GetCurrentPath() + "/ip/ip.txt"
	ipFile, err := os.Open(path)
	if err != nil {
		log.Printf("未找到ip文件：%s \n", path)
		return
	}
	defer ipFile.Close()
	//按行读ip
	buf := bufio.NewReader(ipFile)
	ipWg := &sync.WaitGroup{}
	end := false
	for {
		ip, err := buf.ReadString('\n')
		ip = strings.TrimSpace(ip)
		if err != nil {
			if err == io.EOF {
				end = true
				break
			} else {
				return
			}
		}
		//ip验证
		ipWg.Add(1)
		go func(ip string) {
			proxyTest(ip, ipWg)
		}("http://" + ip)
	}
	ipWg.Wait()
	for !end {
	}
	close(ipCh)
	for ip := range ipCh {
		ipArr = append(ipArr, ip)
	}
	ipArr = append(ipArr, "")
	fmt.Printf("共找到 %d 个 可用 ip\n", len(ipArr)-1)
	return
}
