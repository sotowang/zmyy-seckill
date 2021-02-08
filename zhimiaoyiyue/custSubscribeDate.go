package zhimiaoyiyue

import (
	"fmt"
	"strconv"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/model"
	"zmyy_seckill/util"
)

/**
/sc/wx/HandlerSubscribe.ashx?act=GetCustSubscribeDateAll&pid=2&id=1921&month=202102
*/
func (e *ZMYYEngine) GetCustSubscribeDateAll(customerId, productId, month int) (*model.SubscribeDate, error) {
	url := consts.CustSubscribeDateUrl + "&pid=" + strconv.Itoa(productId) + "&id=" + strconv.Itoa(customerId) + "&month=" + strconv.Itoa(month)
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["cookie"] = e.Conf.Cookie
	bytes, err := fetcher.Fetch(url, headers)
	if err != nil {
		fmt.Printf("GetCustomerProduct() err : %v \n", err)
	}
	subsDates := model.SubscribeDate{}
	err = util.Transfer2SubscribeDateModel(bytes, &subsDates)
	if err != nil {
		fmt.Printf("GetCustSubscribeDateAll() err: %v\n ", err)
		return nil, err
	}
	return &subsDates, nil
}
func (e *ZMYYEngine) GetCustSubscribeDateDetail(date string, productId, customerId int) (*model.SubscribeDateDetail, error) {
	url := consts.CustSubscribeDateDetailUrl + "&pid=" + strconv.Itoa(productId) + "&id=" + strconv.Itoa(customerId) + "&scdate=" + date
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["cookie"] = e.Conf.Cookie
	zftsl, _ := util.GetZFTSL()
	headers["zftsl"] = zftsl
	bytes, err := fetcher.Fetch(url, headers)
	if err != nil {
		fmt.Printf("GetCustSubscribeDateDetail() err : %v \n", err)
	}
	dateDetails := &model.SubscribeDateDetail{}
	err = util.Transfer2SubscribeDateDetailModel(bytes, dateDetails)
	if err != nil {
		fmt.Printf("GetCustSubscribeDateDetail() err: %v\n ", err)
		return nil, err
	}
	dateDetails.Date = date
	return dateDetails, nil
}
