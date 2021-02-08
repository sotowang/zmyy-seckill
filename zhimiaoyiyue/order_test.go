package zhimiaoyiyue

import (
	"testing"
)

func init() {
	e.Init()
}
func TestSaveOrder(t *testing.T) {
	e.GetVerifyPic("2021-02-27", "2", "")
	guid, err := e.CaptchaVerify("2021-02-27", "2", "")
	_, err = e.SaveOrder("2021-02-27", "2", guid, "AAAAAOJdAAAzYjQB")
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestZMYYEngine_GetOrderStatus(t *testing.T) {
	e.GetOrderStatus()
}
