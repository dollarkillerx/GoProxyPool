package reptile

import (
	"GoProxyPool/config"
	"log"
	"time"
)

// 启动爬虫
func RunReptile() {
	log.Println("====================+")
	go reptileAbroad() // 爬取国外

	go reptileDomestic() // 爬取国内
}

// 动态更新数据
func TimingRep() {
	for {
		select {
		// 定时任务
		case <-time.After(config.MyConfig.App.Time * time.Hour):
			// 启动一个协程 去爬取
			// 启动一个协程 吧目标站点弄伤了
			go RunReptile()
		}
	}
}
