package config

import (
	"fmt"
	"testing"
)

func TestGetconfig(t *testing.T) {
	var yaml RootConf
	conf, err := yaml.GetConf()
	if err != nil {
		t.Errorf("failed ")
	}
	fmt.Printf("%v \n", conf)

}
