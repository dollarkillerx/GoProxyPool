package middleware

import (
	"GoProxyPool/config"
	"github.com/gin-gonic/gin"
)

// 限流
func Limiting(ctx *gin.Context) {
	ints := make(chan int, config.MyConfig.App.MaxRequest)

	if len(ints) <= config.MyConfig.App.MaxRequest {
		ints <- 1
		ctx.Next()
		<-ints
		return
	} else {
		ctx.JSON(411, "流量")
		ctx.Abort()
	}

}
