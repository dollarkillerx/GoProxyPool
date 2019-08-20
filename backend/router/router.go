package router

import (
	"GoProxyPool/backend/controller"
	"GoProxyPool/backend/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter() *gin.Engine {
	app := gin.New()

	// 注册
	app.LoadHTMLGlob("view/*")

	router(app)

	return app
}

func router(app *gin.Engine) {

	// 获取国内
	app.GET("/api/domestic", middleware.Limiting, controller.Domestic)

	// 获取国外
	app.GET("/api/foreign", middleware.Limiting, controller.Foreign)

	// 展示页面
	app.GET("/", middleware.Limiting, controller.Home)
}
