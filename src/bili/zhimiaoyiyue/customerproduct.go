package zhimiaoyiyue

import (
	"fmt"
	"strconv"
	"zmyy_seckill/src/bili/consts"
	"zmyy_seckill/src/bili/fetcher"
	"zmyy_seckill/src/bili/model"
	"zmyy_seckill/src/bili/util"
)

func (e *ZMYYEngine) GetCustomerProduct(customerId int) (*model.RootSource, error) {
	url := consts.CustomerProductURL + "&id=" + strconv.Itoa(customerId)
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["cookie"] = "ASP.NET_SessionId=" + consts.SessionId
	bytes, err := fetcher.Fetch(url, headers)
	if err != nil {
		fmt.Printf("GetCustomerProduct() err : %v \n", err)
	}
	customerProducts := model.RootSource{}
	err = util.TransferToCustomerProductListModel(bytes, &customerProducts)
	if err != nil {
		fmt.Printf("GetCustomerProduct() err: %v\n ", err)
		return nil, err
	}
	return &customerProducts, nil

}
