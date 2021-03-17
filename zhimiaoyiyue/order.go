package zhimiaoyiyue

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/model"
	"zmyy_seckill/utils"
)

func (e *ZMYYEngine) SaveOrder(dateDetail model.DateDetail, productId string, guid string) (ok bool, err error) {
	ok = false
	url := consts.SaveUrl + "&birthday=" + e.Conf.Birthday + "&tel=" + e.Conf.Tel + "&sex=" + strconv.Itoa(e.Conf.Sex) + "&cname=" + utils.UrlEncode(e.Conf.Name) + "&doctype=1&idcard=" + e.Conf.IdCard + "&mxid=" + dateDetail.Mxid + "&date=" + dateDetail.Date + "&pid=" + productId + "&Ftime=1&guid=" + guid
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["Cookie"] = e.Conf.Cookie
	headers["Connection"] = "keep-alive"
	zftsl, _ := utils.GetZFTSL()
	headers["zftsl"] = zftsl
	bytes, err := fetcher.Fetch(url, headers)
	if err != nil {
		fmt.Printf("SaveOrder() err : %v \n", err)
		return ok, err
	}
	//如果状态码为200，则订单提交成功
	res := string(bytes)
	if strings.Index(res, `"status":200`) == -1 {
		fmt.Printf("订单 %s: %s-%s 提交失败：%s \n", dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime, res)
		return ok, errors.New("订单提交失败：" + res)
	}
	fmt.Printf("订单 %s: %s-%s 提交成功 \n", dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime)
	//获取订单状态
	ok, _, err = e.GetOrderStatus(dateDetail)
	return ok, err
}

/**
获取订单状态
*/
func (e *ZMYYEngine) GetOrderStatus(dateDetail model.DateDetail) (bool, []byte, error) {
	url := consts.OrderStatusUrl
	headers := make(map[string]string)
	headers["Referer"] = consts.Refer
	zftsl, _ := utils.GetZFTSL()
	headers["zftsl"] = zftsl
	headers["Cookie"] = e.Conf.Cookie
	resp, err := fetcher.Fetch(url, headers)
	if err != nil {
		return false, nil, err
	}
	//如果状态码为200，则订单提交成功
	res := string(resp)
	if strings.Index(res, `"status":200`) == -1 {
		fmt.Printf("订单: %s: %s-%s 提交成功，但实际未生效：%s\n", dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime, res)
		return false, resp, errors.New("订单实际未生效：" + res)
	}
	return true, resp, nil
}

/**
封装下单过程
*/
func (e *ZMYYEngine) Bingo(dateDetail model.DateDetail, productId string) (bool, error) {
	try := 0
	path := ""
	var err error
	//1.获取验证码,重试3次
	fmt.Printf("正在获取验证码图片：%s  %s-%s\n", dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime)
	for path == "" && try < 3 {
		path, err = e.GetVerifyPic(dateDetail)
		try++
	}
	if path == "" {
		return false, err
	}
	//2.识别验证码,重试3次
	try = 0
	guid := ""
	fmt.Printf("正在识别验证码图片：%s  %s-%s\n", dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime)
	for guid == "" && try < 3 {
		guid, err = e.CaptchaVerify(path)
		try++
	}
	if guid == "" {
		return false, err
	}
	//3.下单,重试3次
	fmt.Printf("验证码图片:%s  %s-%s 识别成功，尝试下单中...\n", dateDetail.Date, dateDetail.StartTime, dateDetail.EndTime)
	ok := false
	try = 0
	for !ok && try < 3 {
		ok, err = e.SaveOrder(dateDetail, productId, guid)
		try++
	}
	if ok {
		return true, nil
	}
	return false, err
}
