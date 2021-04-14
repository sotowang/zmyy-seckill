package zhimiaoyiyue

import (
	"log"
	"strconv"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/model"
	"zmyy_seckill/utils"
)

/**
/sc/wx/HandlerSubscribe.ashx?act=GetCustSubscribeDateAll&pid=2&id=1921&month=202102
*/
func (e *ZMYYEngine) GetCustSubscribeDateAll(customerId, productId, month int, ip ...string) *model.SubscribeDate {
	url := consts.CustSubscribeDateUrl + "&pid=" + strconv.Itoa(productId) + "&id=" + strconv.Itoa(customerId) + "&month=" + strconv.Itoa(month)
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["cookie"] = e.Conf.Cookie
	headers["Connection"] = consts.Connection
	bytes, err := fetcher.Fetch(url, headers, ip...)
	if err != nil {
		log.Printf("GetCustSubscribeDateAll() err : %v \n", err)
		return nil
	}
	subsDates := model.SubscribeDate{}
	err = utils.Transfer2SubscribeDateModel(bytes, &subsDates)
	if err != nil {
		log.Printf("GetCustSubscribeDateAll() err: %v\n ", err)
		return nil
	}
	return &subsDates
}
func (e *ZMYYEngine) GetCustSubscribeDateDetail(date string, productId, customerId int, ip ...string) *model.SubscribeDateDetail {
	url := consts.CustSubscribeDateDetailUrl + "&pid=" + strconv.Itoa(productId) + "&id=" + strconv.Itoa(customerId) + "&scdate=" + date
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["cookie"] = e.Conf.Cookie
	zftsl := utils.GetZFTSL()
	headers["zftsl"] = zftsl
	bytes, err := fetcher.FetchWithRatelimter(url, headers, ip...)
	if err != nil {
		return nil
	}
	dateDetails := &model.SubscribeDateDetail{}
	err = utils.Transfer2SubscribeDateDetailModel(bytes, dateDetails)
	if err != nil {
		return nil
	}
	dateDetails.Date = date
	return dateDetails
}
