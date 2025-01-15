package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 更新
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	})
	r.Run()
}
