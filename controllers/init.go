package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/universalmacro/common/server"
	"github.com/universalmacro/core/services"
	"github.com/universalmacro/core/services/models"
)

var router = gin.Default()
var VERSION = "0.0.3"

type Headers struct {
	Authorization string
	ApiKey        *string
}

func Init(addr ...string) {
	var adminController = newAdminController()
	var adminService = services.GetAdminService()
	var sessionsControllers = newSessionsController()
	var nodeController = newNodeController()
	// Cors
	router.Use(server.CorsMiddleware())
	// Auth
	router.Use(func(ctx *gin.Context) {
		var headers Headers
		ctx.ShouldBindHeader(&headers)
		authorization := headers.Authorization
		splited := strings.Split(authorization, " ")
		if authorization != "" && len(splited) == 2 {
			admin, err := adminService.VerifyToken(splited[1])
			if admin != nil && err == nil {
				ctx.Set("admin", admin)
			}
		}
		ctx.Next()
	})
	server.MetricsMiddleware(router)
	router.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version": VERSION,
		})
	})
	// Session
	router.POST("/sessions", sessionsControllers.CreateSession)

	// Admin
	router.GET("/admins", adminController.ListAdmin)
	router.POST("/admins", adminController.CreateAdmin)
	router.GET("/admins/self", adminController.GetSelf)
	router.PUT("/admins/self/password", adminController.UpdateSelfPassword)
	router.PUT("/admins/:id/password", adminController.UpdatePassword)
	router.DELETE("/admins/:id", adminController.DeleteAdmin)
	router.GET("/admins/:id", adminController.GetAdmin)

	// Node
	router.POST("/nodes", nodeController.CreateNode)
	router.GET("/nodes", nodeController.ListNode)
	router.GET("/nodes/:id/config/database", nodeController.GetNodeDatabaseConfig)
	router.PUT("/nodes/:id/config/database", nodeController.UpdateNodeDatabaseConfig)
	router.GET("/nodes/:id/config/redis", nodeController.GetNodeRedisConfig)
	router.PUT("/nodes/:id/config/redis", nodeController.UpdateNodeRedisConfig)
	router.DELETE("/nodes/:id", nodeController.DeleteNode)
	router.GET("/nodes/:id", nodeController.GetNode)
	router.Run(addr...)
}

func getAdmin(ctx *gin.Context) *models.Admin {
	adminInterface, ok := ctx.Get("admin")
	if !ok {
		return nil
	}
	admin, ok := adminInterface.(*models.Admin)
	if !ok {
		return nil
	}
	return admin
}
