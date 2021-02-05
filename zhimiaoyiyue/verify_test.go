package zhimiaoyiyue

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
	"time"
)

var e = ZMYYEngine{}

func init() {
	e.Init()
}
func TestSave20(t *testing.T) {
	e.Save20("2021-02-05")
}
func TestZMYYEngine_CaptchaVerifyUrl(t *testing.T) {
	timeUnix := time.Now().Unix()
	md5 := md5.New()
	md5.Write([]byte(strconv.FormatInt(timeUnix, 10)))
	str := md5.Sum(nil)

	pic, err := e.GetVerifyPic(hex.EncodeToString(str))
	if err != nil {
		t.Errorf("err : %v", err)
	}
	fmt.Printf("%v", pic)
}
