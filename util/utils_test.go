package util

import (
	"fmt"
	"testing"
)

func TestBase64ToPics(t *testing.T) {
	Base64ToPics()
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
