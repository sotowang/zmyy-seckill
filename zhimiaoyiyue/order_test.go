package zhimiaoyiyue

import (
	"testing"
)

func init() {
	e.Init()
}
func TestSaveOrder(t *testing.T) {
	_, err := e.SaveOrder("2021-02-10", "54")
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestZMYYEngine_GetOrderStatus(t *testing.T) {
	e.GetOrderStatus()
}
