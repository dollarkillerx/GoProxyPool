package router

import "github.com/gin-gonic/gin"

func RegisterRouter() *gin.Engine {
	app := gin.New()

	router(app)

	return app
}

func router(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200,"asdas")
	})
}