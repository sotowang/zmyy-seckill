package ip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
	"zmyy_seckill/consts"
)

type IPModel struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}
type Data struct {
	CurrentPage int      `json:"current_page"`
	IPs         []IPInfo `json:"data"`
	LastPage    string   `json:"last_page"`
	PerPage     string   `json:"per_page"`
	To          int      `json:"to"`
	Total       int      `json:"total"`
}
type IPInfo struct {
	IP          string `json:"ip"`
	Port        string `json:"port"`
	IpAddress   string `json:"ip_address"`
	Speed       int    `json:"speed"`
	Anonymity   int    `json:"anonymity"`
	ISP         string `json:"isp"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
	UniqueId    string `json:"unique_id"`
	ValidatedAt string `json:"validated_at"`
	Protocol    string `json:"protocol"`
}

//ip获取
var ip_urls = [...]string{"https://ip.jiangxianli.com/api/proxy_ips"}
var wg sync.WaitGroup

func fectch(url string) ([]string, error) {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return filterIP(contents)
}

/**
过滤代理ip
*/
func filterIP(contents []byte) ([]string, error) {
	ipModel := IPModel{}
	json.Unmarshal(contents, &ipModel)
	var ipArr []string
	ch := make(chan string, len(ipModel.Data.IPs))
	for _, v := range ipModel.Data.IPs {
		wg.Add(1)
		go func(ip string) {
			proxyTest(ip, ch)
		}("http://" + v.IP + ":" + v.Port)
	}
	go func(in <-chan string) {
		for v := range in {
			ipArr = append(ipArr, v)
		}
	}(ch)
	wg.Wait()
	return ipArr, nil
}

/**
判断代理ip的有效性
*/
func proxyTest(ip string, out chan<- string) {
	defer wg.Done()
	testUrl := "http://www.baidu.com"
	// 解析代理地址
	proxy, err := url.Parse(ip)
	//设置网络传输
	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(3),
	}
	// 创建连接客户端
	httpClient := &http.Client{
		Timeout:   time.Second * 3,
		Transport: netTransport,
	}
	res, err := httpClient.Get(testUrl)

	if err != nil || res.StatusCode != http.StatusOK {
		return
	}
	defer res.Body.Close()
	fmt.Printf("有效ip : %v \n", ip)
	out <- ip
}
func FetchIp() (ip []string) {
	fmt.Printf("尝试获取代理ip...\n")
	for _, url := range ip_urls {
		ip, _ = fectch(url)
		fmt.Printf("共找到 %d 个 可用 ip\n", len(ip))
		return
	}
	return nil
}
func SetRandomIP() string {
	if consts.ProxyIpArr == nil || len(consts.ProxyIpArr) == 0 {
		//如果IP池IP用尽，则使用本机IP
		consts.ProxyIp = ""
		return consts.ProxyIp
	}
	//随机从ip池中获取一个ip
	index := rand.Intn(len(consts.ProxyIpArr))
	consts.ProxyIp = consts.ProxyIpArr[index]
	//删除该ip
	consts.ProxyIpArr = append(consts.ProxyIpArr[:index], consts.ProxyIpArr[index+1:]...)
	return consts.ProxyIp
}
