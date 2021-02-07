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
	"strings"
	"zmyy_seckill/model"
)

func Transfer2CustomerListModel(jsonCont []byte, cumtomers *model.CustomerList) error {
	err := json.Unmarshal(jsonCont, &cumtomers)
	if err != nil {
		return err
	}
	return nil
}
func Transfer2CustomerProductListModel(jsonCont []byte, m *model.RootSource) error {
	err := json.Unmarshal(jsonCont, &m)
	if err != nil {
		return err
	}
	return nil
}

func Transfer2SubscribeDateModel(jsonCont []byte, m *model.SubscribeDate) error {
	err := json.Unmarshal(jsonCont, &m)
	if err != nil {
		return err
	}
	return nil
}
func Transfer2VerifyResultModel(jsonCont []byte, m *model.VerifyResultModel) error {
	err := json.Unmarshal(jsonCont, &m)
	if err != nil {
		return err
	}
	return nil
}

//将Base64文件（../imgs/veryfiPics）转成图片
func Base64ToPics() error {
	data, err := ioutil.ReadFile("../imgs/verifyPics")
	if err != nil {
		fmt.Printf("can not load file err : %v\n", err)
		return err
	}
	m := &model.VerifyPicModel{}
	err = json.Unmarshal(data, m)
	if err != nil {
		return err
	}
	d, _ := base64.StdEncoding.DecodeString(m.Dragon)
	t, _ := base64.StdEncoding.DecodeString(m.Tiger)
	fd, _ := os.OpenFile("../imgs/dragon.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	ft, _ := os.OpenFile("../imgs/tiger.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer fd.Close()
	defer ft.Close()
	_, err = fd.Write(d)
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

/**
  调用Python脚本破解验证码
*/
func CallPythonScript(tigerPath, dragonPath, procssPath string) (string, error) {
	exePath := "../pyexe/main/main.exe"
	args := []string{tigerPath, dragonPath, procssPath}
	out, err := exec.Command(exePath, args...).Output()
	if err != nil {
		fmt.Printf("CallPythonScript err: %v\n", err)
		return "", err
	}
	str := strings.Replace(string(out), "\r", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	fmt.Printf("滑块坐标为： %s\n", str)
	return str, nil
}

func GetZFTSL() (string, error) {
	bytes, err := ioutil.ReadFile("../js/app.js")
	if err != nil {
		fmt.Printf("GetZFTSL err : %v\n", err)
		return "", err
	}
	vm := otto.New()
	_, err = vm.Run(string(bytes))
	enc, err := vm.Call("zftsl", nil)
	if err != nil {
		fmt.Printf("GetZFTSL err : %v\n", err)
		return "", err
	}
	fmt.Printf("zftsl : %s\n", enc.String())
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
