/**
* Created by GoLand
* User: dollarkiller
* Date: 19-7-13
* Time: 上午8:41
* */
package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type myConf struct {
	App struct {
		Time       time.Duration `yaml:"time"`
		MaxRequest int           `yaml:"max_request"`
		Port       string        `yaml:"port"`
	}
}

var (
	MyConfig *myConf
)

func init() {
	MyConfig = &myConf{}

	bytes, e := ioutil.ReadFile("./config.yml")
	if e != nil {
		panic(e.Error())
	}

	e = yaml.Unmarshal(bytes, MyConfig)
	if e != nil {
		panic(e.Error())
	}

}
