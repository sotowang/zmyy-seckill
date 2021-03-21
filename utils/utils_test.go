package utils

import (
	"fmt"
	"testing"
)

func TestBase64ToPics(t *testing.T) {
	prefix := "2021-02-18:08:00:00-09:00:00-2"
	Base64ToPics(prefix)
}
func TestCallPythonScript(t *testing.T) {
	dragonPath := "C:\\Users\\Administrator\\IdeaProjects\\zmyy_seckill\\imgs\\2021-03-23-08_00_00-11_30_00-dragon.png"
	tigerPath := "C:\\Users\\Administrator\\IdeaProjects\\zmyy_seckill\\imgs\\2021-03-23-08_00_00-11_30_00-tiger.png"
	processPath := "C:\\Users\\Administrator\\IdeaProjects\\zmyy_seckill\\imgs\\2021-03-23-08_00_00-11_30_00-process.png"
	pythonScript, _ := CallPythonScript(tigerPath, dragonPath, processPath)
	fmt.Printf("%s", pythonScript)
}
func TestCallJsScript(t *testing.T) {
	zftsl := GetZFTSL()
	fmt.Printf("%v \n", zftsl)
}
