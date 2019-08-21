package controller

import (
	"GoProxyPool/detasource/memory_conn"
	"github.com/gin-gonic/gin"
)

func Domestic(ctx *gin.Context) {
	value, ok := memory_conn.MemoryDb.Load("gn")
	if ok {
		ctx.Header("contentType", "application/json")
		s,ok := value.(string)
		if ok {
			ctx.String(200, s)
		}else {
			ctx.JSON(200,value)
		}
	} else {
		value2, b := memory_conn.MemoryDb.Load(1)
		if b {
			ctx.Header("contentType", "application/json")
			s,ok := value2.(string)
			if ok {
				ctx.String(200, s)
			}else {
				ctx.JSON(200,value2)
			}
		}else {
			ctx.JSON(200,"none")
		}
	}
}

func Foreign(ctx *gin.Context) {
	value, ok := memory_conn.MemoryDb.Load("gw")
	if ok {
		ctx.Header("contentType", "application/json")
		s,ok := value.(string)
		if ok {
			ctx.String(200, s)
		}else {
			ctx.JSON(200,value)
		}
	} else {
		value2, b := memory_conn.MemoryDb.Load(1)
		if b {
			ctx.Header("contentType", "application/json")
			s,ok := value2.(string)
			if ok {
				ctx.String(200, s)
			}else {
				ctx.JSON(200,value2)
			}

		}else {
			ctx.JSON(200,"none")
		}
	}
}
