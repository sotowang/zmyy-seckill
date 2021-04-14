package zhimiaoyiyue

import (
	"fmt"
	"testing"
)

func TestZMYYEngine_GetCustSubscribeDateAll(t *testing.T) {
	e := ZMYYEngine{}
	all := e.GetCustSubscribeDateAll(1921, 1, 202102)
	fmt.Printf("%v", all)
}
func TestZMYYEngine_GetCustSubscribeDateDetail(t *testing.T) {
	e := ZMYYEngine{}
	all, err := e.GetCustSubscribeDateDetail("2021-02-27", 2, 1921)
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("%v", all)
}
