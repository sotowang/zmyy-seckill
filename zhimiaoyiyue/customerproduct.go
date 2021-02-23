package zhimiaoyiyue

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/model"
	"zmyy_seckill/util"
)

//获取换购疫苗的productId
func (e *ZMYYEngine) GetCustomerProduct(customerId int) (int, error) {
	url := consts.CustomerProductURL + "&id=" + strconv.Itoa(customerId) + "&lat=33.54256&lng=114.61165"
	url = "https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx?act=CustomerProduct&id=1921&lat=33.54256&lng=114.61165"
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	headers["cookie"] = e.Conf.Cookie
	headers["Connection"] = consts.Connection
	headers["Host"] = "cloud.cn2030.com"
	headers["content-type"] = "application/json"
	zftsl, _ := util.GetZFTSL()
	headers["zftsl"] = zftsl

	bytes, err := fetcher.Fetch(url, headers)
	if err != nil {
		log.Printf("GetCustomerProduct() err : %v \n", err)
		return -1, err
	}
	customerProducts := model.RootSource{}
	err = util.Transfer2CustomerProductListModel(bytes, &customerProducts)
	if err != nil {
		log.Printf("GetCustomerProduct() err: %v\n ", err)
		return -1, err
	}
	fmt.Printf("正在查找疫苗信息：\n")
	var selected model.CustomerProduct
	var contains bool
	for k, v := range customerProducts.CustomerProducts {
		fmt.Printf("第 %d个疫苗：%s productId: %d\n", k+1, v.Text, v.Id)
		if v.Text == e.Conf.ProductName {
			selected = v
			contains = true
		}
	}
	if contains {
		fmt.Printf("选中疫苗：%s，其productId为 %d\n", selected.Text, selected.Id)
		return selected.Id, nil
	}

	return -1, errors.New("未找到指定疫苗，请对比配置文件疫苗信息是否正确！")

}
