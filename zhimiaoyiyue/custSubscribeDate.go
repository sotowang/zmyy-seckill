package zhimiaoyiyue

import (
	"fmt"
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
func (e *ZMYYEngine) GetCustSubscribeDateAll(customerId, productId, month int) (*model.SubscribeDate, error) {
	url := consts.CustSubscribeDateUrl + "&pid=" + strconv.Itoa(productId) + "&id=" + strconv.Itoa(customerId) + "&month=" + strconv.Itoa(month)
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["cookie"] = e.Conf.Cookie
	headers["Connection"] = consts.Connection
	bytes, err := fetcher.Fetch(url, headers)
	if err != nil {
		fmt.Printf("GetCustSubscribeDateAll() err : %v \n", err)
		return nil, err
	}
	subsDates := model.SubscribeDate{}
	err = utils.Transfer2SubscribeDateModel(bytes, &subsDates)
	if err != nil {
		fmt.Printf("GetCustSubscribeDateAll() err: %v\n ", err)
		return nil, err
	}

	if len(subsDates.Dates) == 0 {
		fmt.Printf("GetCustSubscribeDateAll:未找到可预约的日期，将重新查找...\n")
	} else {
		cnt := 0
		for index, date := range subsDates.Dates {
			if date.Enable {
				cnt++
				fmt.Printf("日期%d： %s\n", index+1, date.Date)
			}
		}
		if cnt == 0 {
			fmt.Printf("未找到可预约的日期，将重新查找...\n")
		} else {
			fmt.Printf("共找到 %d 可预约的日期\n", cnt)
		}

	}
	return &subsDates, nil
}
func (e *ZMYYEngine) GetCustSubscribeDateDetail(date string, productId, customerId int) (*model.SubscribeDateDetail, error) {
	url := consts.CustSubscribeDateDetailUrl + "&pid=" + strconv.Itoa(productId) + "&id=" + strconv.Itoa(customerId) + "&scdate=" + date
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["cookie"] = e.Conf.Cookie
	zftsl, _ := utils.GetZFTSL()
	headers["zftsl"] = zftsl
	bytes, err := fetcher.Fetch(url, headers)
	if err != nil {
		log.Printf("GetCustSubscribeDateDetail() err : %v \n", err)
	}
	dateDetails := &model.SubscribeDateDetail{}
	err = utils.Transfer2SubscribeDateDetailModel(bytes, dateDetails)
	if err != nil {
		log.Printf("GetCustSubscribeDateDetail() err: %v\n ", err)
		return nil, err
	}
	dateDetails.Date = date
	return dateDetails, nil
}
