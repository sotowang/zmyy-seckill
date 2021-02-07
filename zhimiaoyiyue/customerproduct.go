package zhimiaoyiyue

import (
	"fmt"
	"strconv"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/model"
	"zmyy_seckill/util"
)

func (e *ZMYYEngine) GetCustomerProduct(customerId int) (*model.RootSource, error) {
	url := consts.CustomerProductURL + "&id=" + strconv.Itoa(customerId)
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["cookie"] = e.Conf.Cookie
	bytes, err := fetcher.Fetch(url, headers)
	if err != nil {
		fmt.Printf("GetCustomerProduct() err : %v \n", err)
	}
	customerProducts := model.RootSource{}
	err = util.Transfer2CustomerProductListModel(bytes, &customerProducts)
	if err != nil {
		fmt.Printf("GetCustomerProduct() err: %v\n ", err)
		return nil, err
	}
	return &customerProducts, nil

}
