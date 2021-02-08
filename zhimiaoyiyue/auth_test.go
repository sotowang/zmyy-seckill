package zhimiaoyiyue

import (
	"fmt"
	"testing"
	"zmyy_seckill/util"
)

func TestAuth(t *testing.T) {
	e := ZMYYEngine{}
	e.Init()
	e.AuthAndSetSessionID()
}

func TestParseSessionId(t *testing.T) {
	id := util.ParseSessionId("ASP.NET_SessionId=jw1c3itgmqxoik0q3sazbyx5; path=/; HttpOnly; SameSite=Lax")
	fmt.Printf("%s", id)
}
