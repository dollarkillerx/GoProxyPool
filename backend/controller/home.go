package controller

import "github.com/gin-gonic/gin"

func Home(ctx *gin.Context) {
	ctx.HTML(200,"home.html",nil)
}
