package zhimiaoyiyue

import (
	"fmt"
	"testing"
)

func TestZMYYEngine_GetCustSubscribeDateAll(t *testing.T) {
	e := ZMYYEngine{}
	all, err := e.GetCustSubscribeDateAll(1921, 1, 202102)
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Printf("%v", all)
}
