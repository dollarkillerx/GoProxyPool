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
	"log"
	"time"
)

type myconf struct {
	App struct {
		Host       string `yaml:"host"`
		BaseUrl    string `yaml:"baseurl"`
		Debug      bool   `yaml:"debug"`
		BaseStatic string `yaml:"-"`
		Language   string `yaml:"language"`
		S3url string `yaml:""`
	}
	Mysql struct {
		Dsn   string `yaml:"dsn"`
		Cache bool   `yaml:"cache"`
	}
	Redis struct {
		Maxidle     int           `yaml:"maxidle"`
		MaxActive   int           `yaml:"max_active"`
		IdleTimeout time.Duration `yaml:"idle_timeout"`
		Port        string        `yaml:"port"`
	}
}

var (
	MyConfig *myconf
)

func init() {
	MyConfig = &myconf{}

	bytes, e := ioutil.ReadFile("./config.yml")
	if e != nil {
		panic(e.Error())
	}

	e = yaml.Unmarshal(bytes, MyConfig)
	if e != nil {
		panic(e.Error())
	}

	MyConfig.App.BaseStatic = MyConfig.App.BaseUrl + "/static"

	if MyConfig.App.Debug {
		log.Println(MyConfig)
	}
}
