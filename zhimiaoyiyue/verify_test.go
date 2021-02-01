package zhimiaoyiyue

import "testing"

func TestSave20(t *testing.T) {
	e := ZMYYEngine{}
	e.Init()
	e.Save20("2021-02-05")
}
