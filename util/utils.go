package util

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"zmyy_seckill/model"
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

func TransferToSubscribeDateModel(jsonCont []byte, m *model.SubscribeDate) error {
	err := json.Unmarshal(jsonCont, &m)
	if err != nil {
		return err
	}
	return nil
}

func TransferToVerifyModel(jsonCont []byte, m *model.VerifyPicModel) error {
	err := json.Unmarshal(jsonCont, &m)
	if err != nil {
		return err
	}
	return nil
}

func Base64ToPics(m model.VerifyPicModel) {
	dragon := m.Dragon
	tiger := m.Tiger
	d, _ := base64.StdEncoding.DecodeString(dragon)
	t, _ := base64.StdEncoding.DecodeString(tiger)
	fd, _ := os.OpenFile("C:\\Users\\Administrator\\IdeaProjects\\zmyy_seckill\\imgs\\dragon.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	ft, _ := os.OpenFile("C:\\Users\\Administrator\\IdeaProjects\\zmyy_seckill\\imgs\\tiger.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer fd.Close()
	defer ft.Close()
	fd.Write(d)
	ft.Write(t)
}

func CallPythonScript(tigerPath, dragonPath string) []byte {
	args := []string{"C:\\Users\\Administrator\\IdeaProjects\\SlideCrack\\slide_01\\main.py", tigerPath, dragonPath}
	out, err := exec.Command("python", args...).Output()
	if err != nil {
		return nil
	}
	return out
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
