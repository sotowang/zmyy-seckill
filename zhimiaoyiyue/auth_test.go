package zhimiaoyiyue

import (
	"fmt"
	"testing"
	"zmyy_seckill/util"
)

//[]unit8 [247,104,203,177,74,198,145,76,58,185,13,112]
//subjectKeyId [143,234,134,236,117,103,43,186,84,126,14,24,249,248,0,63,198,109,36,61]
//authorKeyID [127,211,153,243,160,71,14,49,0,86,86,34,142,183,204,158,221,202,1,138]
func TestAuth(t *testing.T) {

	zftsl := "b1b7c103702cf5d0e7bf4a1eea80e6de"
	aa := "c0fd75befa2c82babae8a2b6188fd20e"
	//zftsl =  "92jLsUrGkUw6uQ1w"
	//ar := []uint8{247,104,203,177,74,198,145,76,58,185,13,112}
	//subjectKeyId := []uint8{143,234,134,236,117,103,43,186,84,126,14,24,249,248,0,63,198,109,36,61}
	//authorKeyID := []uint8{127,211,153,243,160,71,14,49,0,86,86,34,142,183,204,158,221,202,1,138}
	//encode := base64.StdEncoding.EncodeToString(authorKeyID)
	fmt.Printf("%d\n", len(zftsl))
	fmt.Printf("%d\n", len(aa))
	//AuthAndSetSessionID()
}

func TestParseSessionId(t *testing.T) {
	id := util.ParseSessionId("ASP.NET_SessionId=jw1c3itgmqxoik0q3sazbyx5; path=/; HttpOnly; SameSite=Lax")
	fmt.Printf("%s", id)
}
