package zhimiaoyiyue

import (
	"errors"
	"fmt"
	"strconv"
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/model"
	"zmyy_seckill/util"
)

//获取换购疫苗的productId
func (e *ZMYYEngine) GetCustomerProduct(customerId int) (int, error) {
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
		return -1, err
	}
	fmt.Printf("正在查找疫苗信息：\n")
	for k, v := range customerProducts.CustomerProducts {
		fmt.Printf("第 %d个疫苗：%s\n", k+1, v.Text)
		if v.Text == e.Conf.ProductName {
			fmt.Printf("选中第 %d个疫苗：%s，其productId为 %d\n", k+1, v.Text, v.Id)
			return v.Id, nil
		}

	}
	return -1, errors.New("未找到指定疫苗，请对比配置文件疫苗信息是否正确！")

}
