package zhimiaoyiyue

import (
	"fmt"
	"testing"
)

var e = ZMYYEngine{}

func init() {
	e.Init()
}
func TestSave20(t *testing.T) {
	e.Save20("2021-02-05")
}
func TestZMYYEngine_CaptchaVerify(t *testing.T) {
	m, err := e.CaptchaVerify()
	if err != nil {
		t.Errorf("err: %v\n", err)
	}
	fmt.Printf("%v", m)
}
func TestZMYYEngine_GetVerifyPic(t *testing.T) {
	err := e.GetVerifyPic()
	if err != nil {
		t.Errorf("err : %v", err)
		return
	}
}
