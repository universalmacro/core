package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/universalmacro/common/server"
)

var router = gin.Default()

var version = "0.0.2"

func Init(addr ...string) {
	// var adminController = newAdminController()
	var sessionsControllers = newSessionsController()
	router.Use(server.CorsMiddleware())
	server.MetricsMiddleware(router)
	router.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version": version,
		})
	})
	router.POST("/sessions", sessionsControllers.CreateSession)
	router.Run(addr...)
}
