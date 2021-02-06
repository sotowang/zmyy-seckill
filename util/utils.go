package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	"io/ioutil"
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

func Base64ToPics(m model.VerifyPicModel) error {
	dragon := m.Dragon
	tiger := m.Tiger
	d, _ := base64.StdEncoding.DecodeString(dragon)
	t, _ := base64.StdEncoding.DecodeString(tiger)
	fd, _ := os.OpenFile("../imgs/dragon.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	ft, _ := os.OpenFile("../imgs/tiger.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer fd.Close()
	defer ft.Close()
	_, err := fd.Write(d)
	if err != nil {
		fmt.Printf("Base64ToPics() err : %v", err)
		return err
	}
	_, err = ft.Write(t)
	if err != nil {
		fmt.Printf("Base64ToPics() err : %v", err)
		return err
	}
	return nil
}

func CallPythonScript(tigerPath, dragonPath, procssPath string) ([]byte, error) {
	exePath := "../pyexe/main/main.exe"
	args := []string{tigerPath, dragonPath, procssPath}
	out, err := exec.Command(exePath, args...).Output()
	if err != nil {
		fmt.Printf("CallPythonScript err: %v\n", err)
		return nil, err
	}
	return out, nil
}

func CallJsScript() (string, error) {
	bytes, err := ioutil.ReadFile("../js/app.js")
	if err != nil {
		fmt.Printf("CallJsScript err : %v\n", err)
		return "", err
	}
	vm := otto.New()
	_, err = vm.Run(string(bytes))
	enc, err := vm.Call("zftsl", nil)
	if err != nil {
		fmt.Printf("CallJsScript err : %v\n", err)
		return "", err
	}
	return enc.String(), nil
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
