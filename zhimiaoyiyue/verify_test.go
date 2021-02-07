package zhimiaoyiyue

import (
	"fmt"
	"testing"
)

var e = ZMYYEngine{}

func init() {
	e.Init()
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
