package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/universalmacro/common/server"
	coreapiinterfaces "github.com/universalmacro/core-api-interfaces"
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
	// var merchantController = newMerchantController()
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
	coreapiinterfaces.AdminApiBinding(router, adminController)
	coreapiinterfaces.NodeApiBinding(router, nodeController)
	coreapiinterfaces.SessionApiBinding(router, sessionsControllers)
	// coreapiinterfaces.MerchantApiBinding(router, merchantController)
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
