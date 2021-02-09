package zhimiaoyiyue

import (
	"testing"
)

func init() {
	e.Init()
}

func TestSaveOrder(t *testing.T) {
	path, _ := e.GetVerifyPic(dateDetail)
	guid, err := e.CaptchaVerify(path)
	_, err = e.SaveOrder(dateDetail, "2", guid)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestZMYYEngine_GetOrderStatus(t *testing.T) {
	e.GetOrderStatus(dateDetail)
}
