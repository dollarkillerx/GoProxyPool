package router

import (
	"GoProxyPool/backend/controller"
	"GoProxyPool/backend/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter() *gin.Engine {
	app := gin.New()

	router(app)

	return app
}

func router(app *gin.Engine) {

	// 获取国内
	app.GET("/domestic", middleware.Limiting, controller.Domestic)

	// 获取国外
	app.GET("/foreign", middleware.Limiting, controller.Foreign)
}
