package controller

import (
	"GoProxyPool/detasource/memory_conn"
	"github.com/gin-gonic/gin"
)

func Domestic(ctx *gin.Context) {
	value, ok := memory_conn.MemoryDb.Load("gn")
	if ok {
		ctx.Header("contentType", "application/json")
		ctx.String(200, value.(string))
	} else {
		value2, b := memory_conn.MemoryDb.Load(1)
		if b {
			ctx.Header("contentType", "application/json")
			ctx.String(200, value2.(string))
		}
	}
}

func Foreign(ctx *gin.Context) {
	value, ok := memory_conn.MemoryDb.Load("gw")
	if ok {
		ctx.Header("contentType", "application/json")
		ctx.String(200, value.(string))
	} else {
		value2, b := memory_conn.MemoryDb.Load(1)
		if b {
			ctx.Header("contentType", "application/json")
			ctx.String(200, value2.(string))

		}
	}
}
