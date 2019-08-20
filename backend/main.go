package main

import (
	"GoProxyPool/backend/router"
	"GoProxyPool/config"
	"GoProxyPool/reptile"
	"github.com/dollarkillerx/easyutils/clog"
)

func main() {
	app := router.RegisterRouter()

	// 当服务第一次启动的时候执行爬取计划
	go reptile.RunReptile()
	// 定时更新数据库
	go reptile.TimingRep()

	clog.Println("http://0.0.0.0" + config.MyConfig.App.Port)
	err := app.Run(config.MyConfig.App.Port)
	if err != nil {
		clog.Println(err.Error())
	}
}
