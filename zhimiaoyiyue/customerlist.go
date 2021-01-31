package zhimiaoyiyue

import (
	"zmyy_seckill/consts"
	"zmyy_seckill/fetcher"
	"zmyy_seckill/model"
	"zmyy_seckill/util"
)

func (e *ZMYYEngine) GetCustomerList() (*model.CustomerList, error) {
	params := "[\"" + e.Conf.Province + "\",\"" + e.Conf.City + "\",\"" + e.Conf.District + "\"]"
	newUrl := consts.CustomerListUrl + "&city=" + util.UrlEncode(params) + "&id=0&cityCode=610100&product=" + e.Conf.Product
	headers := make(map[string]string)
	headers["User-Agent"] = consts.UserAgent
	headers["Referer"] = consts.Refer
	bytes, err2 := fetcher.Fetch(newUrl, headers)
	if err2 != nil {
		return nil, err2
	}
	customers := model.CustomerList{}
	err2 = util.TransferToCustomerListModel(bytes, &customers)
	if err2 != nil {
		return nil, err2
	}
	return &customers, nil
}
