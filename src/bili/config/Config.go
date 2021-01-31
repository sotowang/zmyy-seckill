package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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
}

func (c *RootConf) GetConf() (CustomerConf, error) {
	yamlFile, err := ioutil.ReadFile("C:\\Users\\Administrator\\IdeaProjects\\learngo\\src\\bili\\config\\conf.yaml")
	if err != nil {
		log.Printf("yaml file get err #%v ", err)
		return CustomerConf{}, err
	}
	var conf = RootConf{}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Printf("failed to unmarshal : %v\n", err)
	}
	return conf.CustomerConf, nil

}
