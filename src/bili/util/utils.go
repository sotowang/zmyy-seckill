package util

import (
	"encoding/json"
	"net/url"
	"regexp"
	"zmyy_seckill/src/bili/model"
)

func TransferToCustomerListModel(jsonCont []byte, cumtomers *model.CustomerList) error {
	err := json.Unmarshal(jsonCont, &cumtomers)
	if err != nil {
		return err
	}
	return nil
}
func TransferToCustomerProductListModel(jsonCont []byte, m *model.RootSource) error {
	err := json.Unmarshal(jsonCont, &m)
	if err != nil {
		return err
	}
	return nil
}
func UrlEncode(orgUrl string) string {
	encodeUrl := url.QueryEscape(orgUrl)
	return encodeUrl
}
func ParseSessionId(s string) string {

	const sessionIdRe = `ASP.NET_SessionId=([^;]+)`
	compile := regexp.MustCompile(sessionIdRe)
	match := compile.FindSubmatch([]byte(s))
	return string(match[1])
}
