package zhimiaoyiyue

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/model"
	"zmyy_seckill/util"
)

func (e *ZMYYEngine) SaveOrder(dateDetail model.DateDetail, productId string, guid string) (ok bool, err error) {
	ok = false
	url := consts.SaveUrl + "&birthday=" + e.Conf.Birthday + "&tel=" + e.Conf.Tel + "&sex=" + strconv.Itoa(e.Conf.Sex) + "&cname=" + util.UrlEncode(e.Conf.Name) + "&doctype=1&idcard=" + e.Conf.IdCard + "&mxid=" + dateDetail.Mxid + "&date=" + dateDetail.Date + "&pid=" + productId + "&Ftime=1&guid=" + guid
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["Cookie"] = e.Conf.Cookie
	headers["Connection"] = "keep-alive"
	zftsl, _ := util.GetZFTSL()
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
	zftsl, _ := util.GetZFTSL()
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
