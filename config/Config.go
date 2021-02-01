package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type RootConf struct {
	CustomerConf `yaml:"zhiyiyuemiao"`
}
type CustomerConf struct {
	Province string `yaml:province,omitempty default:""`
	City     string `yaml:"city,omitempty" default:""`
	District string `yaml:"district,omitempty" default:""`
	//0为九价
	Product    string `yaml:"product,omitempty" default:1`
	CustomerId int    `yaml:"customerId,omitempty" default:1776`
	Birthday   string `yaml:"birthday,omitempty" default:""`
	Tel        string `yaml:"tel,omitempty" default:""`
	Sex        int    `yaml:"sex,omitempty" default:1`
	Name       string `yaml:"name,omitempty" default:""`
	IdCard     string `yaml:"idcard,omitempty" default:""`
	Mxid       string `yaml:"mxid,omitempty" default:""`
	Guid       string `yaml:"guid,omitempty" default:""`
}

func (c *RootConf) GetConf() (CustomerConf, error) {
	yamlFile, err := ioutil.ReadFile(getCurrentPath() + "/conf.yaml")
	fmt.Printf("path : %s \n", getCurrentPath())
	if err != nil {
		fmt.Printf("yaml file get err #%v \n", err)
		return CustomerConf{}, err
	}
	var conf = RootConf{}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Printf("failed to unmarshal : %v\n", err)
	}
	return conf.CustomerConf, nil
}
func getCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Get current process path failed . err : %v \n", err)
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}
