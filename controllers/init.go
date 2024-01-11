package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Init(addr ...string) {
	router.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version": "0.0.2",
		})
	})
	router.Run(addr...)
}
