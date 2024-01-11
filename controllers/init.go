package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

var version = "0.0.1"

func Init(addr ...string) {
	router.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version": version,
		})
	})
	router.Run(addr...)
}
