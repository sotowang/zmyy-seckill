package config

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"zmyy_seckill/utils"
)

type RootConf struct {
	CustomerConf `yaml:"zhiyiyuemiao"`
}

type CustomerConf struct {
	Province      string `yaml:province,omitempty default:""`
	City          string `yaml:"city,omitempty" default:""`
	District      string `yaml:"district,omitempty" default:""`
	ProductName   string `yaml:"productName,omitempty" default:1`
	Birthday      string `yaml:"birthday,omitempty" default:""`
	Tel           string `yaml:"tel,omitempty" default:""`
	Sex           int    `yaml:"sex,omitempty" default:1`
	Name          string `yaml:"name,omitempty" default:""`
	IdCard        string `yaml:"idcard,omitempty" default:""`
	Cookie        string `yaml:"cookie,omitempty" default:""`
	CityCode      string `yaml:"-"`
	CustomerName  string `yaml:"customerName,omitempty" default:""`
	Month         int    `yaml:"month,omitempty" default:202102`
	SubscribeTime string `yaml:"subscribeTime,omitempty" default:""`
}

func (c *RootConf) GetConf() (CustomerConf, error) {
	path := utils.GetCurrentPath()
	confPath := path + "/config/conf.yaml"
	yamlFile, err := ioutil.ReadFile(confPath)
	fmt.Printf("当前执行路径 : %s \n", path)
	if err != nil {
		fmt.Printf("yaml file get err #%v \n", err)
		return CustomerConf{}, err
	}
	var conf = RootConf{}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Printf("failed to unmarshal : %v\n", err)
	}
	getCityCode(path, &conf)
	return conf.CustomerConf, nil
}
func getCityCode(path string, conf *RootConf) string {
	cityJsonPath := path + "/config/city.json"
	jq := gojsonq.New().File(cityJsonPath).From("citycode")
	if conf.City != "" {
		r := jq.Select("children").Where("name", "=", conf.Province).Get().([]interface{})[0].(map[string]interface{})["children"].([]interface{})
		for _, v := range r {
			vmap := v.(map[string]interface{})
			if vmap["name"] == conf.City {
				conf.CityCode = vmap["value"].(string) + "00"
				break
			}
		}
	} else if conf.Province != "" {
		r := jq.Select("value").Where("name", "=", conf.Province).Get().([]interface{})[0].(map[string]interface{})
		conf.CityCode = r["value"].(string) + "0000"
	}
	return ""
}
