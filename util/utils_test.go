package util

import (
	"fmt"
	"testing"
)

func TestBase64ToPics(t *testing.T) {
	prefix := "2021-02-18:08:00:00-09:00:00-2"
	Base64ToPics(prefix)
}
func TestCallPythonScript(t *testing.T) {
	dragonPath := "../imgs/dragon.png"
	tigerPath := "../imgs/tiger.png"
	processPath := "../imgs/process.png"
	pythonScript, _ := CallPythonScript(tigerPath, dragonPath, processPath)
	fmt.Printf("%s", pythonScript)
}
func TestCallJsScript(t *testing.T) {
	GetZFTSL()
}
