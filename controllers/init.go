package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

var version = "0.0.1"

func Init(addr ...string) {
	var adminController = newAdminController()
	router.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version": version,
		})
	})
	router.POST("/login", adminController.Login)
	router.Run(addr...)
}
