package zhimiaoyiyue

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"
	"sync"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/model"
	"zmyy_seckill/utils"
)

var orderLock sync.Mutex

func (e *ZMYYEngine) SaveOrder(dateDetail model.DateDetail, productId string, guid string, ip ...string) (ok bool, err error) {
	ok = false
	url := consts.SaveUrl + "&birthday=" + e.Conf.Birthday + "&tel=" + e.Conf.Tel + "&sex=" + strconv.Itoa(e.Conf.Sex) + "&cname=" + utils.UrlEncode(e.Conf.Name) + "&doctype=1&idcard=" + e.Conf.IdCard + "&mxid=" + dateDetail.Mxid + "&date=" + dateDetail.Date + "&pid=" + productId + "&Ftime=1&guid=" + guid
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["Cookie"] = e.Conf.Cookie
	headers["Connection"] = "keep-alive"
	zftsl := utils.GetZFTSL()
	headers["zftsl"] = zftsl
	bytes, err := fetcher.FetchWithRatelimter(url, headers, ip...)
	if err != nil {
		log.Printf("%v SaveOrder() err : %v \n", ip, err)
		return ok, err
	}
	//如果状态码为200，则订单提交成功
	res := string(bytes)
	if strings.Index(res, `"status":200`) == -1 {
		log.Printf("订单 %s: %s-%s 提交失败：%s \n", dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime, res)
		return ok, errors.New("订单提交失败：" + res)
	}
	log.Printf("%v 订单 %s: %s-%s 提交成功 \n", ip, dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime)
	//获取订单状态
	ok, _, err = e.GetOrderStatus(dateDetail, ip...)
	return ok, err
}

/**
获取订单状态
*/
func (e *ZMYYEngine) GetOrderStatus(dateDetail model.DateDetail, ip ...string) (bool, []byte, error) {
	url := consts.OrderStatusUrl
	headers := make(map[string]string)
	headers["Referer"] = consts.Refer
	zftsl := utils.GetZFTSL()
	headers["zftsl"] = zftsl
	headers["Cookie"] = e.Conf.Cookie
	resp, err := fetcher.FetchWithRatelimter(url, headers, ip...)
	if err != nil {
		return false, nil, err
	}
	//如果状态码为200，则订单提交成功
	res := string(resp)
	if strings.Index(res, `"status":200`) == -1 {
		log.Printf("%v 订单: %s: %s-%s 提交成功，但实际未生效：%s\n", ip,
			dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime, res)
		return false, resp, errors.New("订单实际未生效：" + res)
	}
	return true, resp, nil
}

/**
封装下单过程
*/
func (e *ZMYYEngine) Bingo(dateDetail model.DateDetail, productId string, ctx context.Context, ip ...string) (bool, error) {
	orderLock.Lock()
	select {
	case <-ctx.Done():
		orderLock.Unlock()
		return false, nil
	default:
		path := ""
		var err error
		//1.获取验证码
		log.Printf("%v 正在获取验证码图片：%s  %s-%s\n", ip, dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime)
		path, err = e.GetVerifyPic(dateDetail, ip...)
		if path == "" {
			orderLock.Unlock()
			return false, err
		}
		//2.识别验证码
		guid := ""
		log.Printf("%v 正在识别验证码图片：%s  %s-%s\n", ip,
			dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime)
		guid, err = e.CaptchaVerify(path, ip...)
		if guid == "" {
			orderLock.Unlock()
			return false, err
		}
		//3.下单,重试2次
		log.Printf("%v 验证码图片:%s  %s-%s 识别成功，尝试下单中...\n", ip,
			dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime)
		ok := false
		try := 0
		for !ok && try < 2 {
			ok, err = e.SaveOrder(dateDetail, productId, guid, ip...)
			try++
		}
		if ok {
			orderLock.Unlock()
			return true, nil
		}
		orderLock.Unlock()
		return false, err
	}
}
