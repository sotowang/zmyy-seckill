package zhimiaoyiyue

import (
	"fmt"
	"testing"
	"zmyy_seckill/model"
)

var e = ZMYYEngine{}
var dateDetail = model.DateDetail{
	Date:      "2021-02-24",
	StartTime: "08:00",
	EndTime:   "11:30",
	Mxid:      "AAAAAM9YAAAwYjQB",
}

func init() {
	e.Init()
}

func TestZMYYEngine_CaptchaVerify(t *testing.T) {
	path, err := e.GetVerifyPic(dateDetail)
	m, err := e.CaptchaVerify(path)
	if err != nil {
		t.Errorf("err: %v\n", err)
	}
	fmt.Printf("guid: %v \n", m)
}
func TestZMYYEngine_GetVerifyPic(t *testing.T) {
	path, err := e.GetVerifyPic(dateDetail)
	if err != nil {
		t.Errorf("err : %v", err)
		return
	}
	fmt.Printf("%s", path)
}
