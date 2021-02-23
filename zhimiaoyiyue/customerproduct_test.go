package zhimiaoyiyue

import (
	"fmt"
	"testing"
)

func TestCustomerProduct(t *testing.T) {
	e := ZMYYEngine{}
	e.Init()
	product, err := e.GetCustomerProduct(1921)
	if err != nil {
		t.Errorf("err : %v", err)
	}
	fmt.Printf("%v", product)
}
