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
	var sessionsControllers = newSessionsController()
	var nodeController = newNodeController()
	router.Use(server.RequestIDMiddleware())
	// Cors
	router.Use(server.CorsMiddleware())
	// Auth
	router.Use(auth())
	server.MetricsMiddleware(router)
	router.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version": VERSION,
		})
	})
	coreapiinterfaces.AdminApiBinding(router, adminController)
	coreapiinterfaces.NodeApiBinding(router, nodeController)
	coreapiinterfaces.SessionApiBinding(router, sessionsControllers)
	coreapiinterfaces.RoleApiBinding(router, newRolesController())
	router.Run(addr...)
}

func auth() func(ctx *gin.Context) {
	var adminService = services.GetAdminService()
	return func(ctx *gin.Context) {
		var headers Headers
		ctx.ShouldBindHeader(&headers)
		authorization := headers.Authorization
		splited := strings.Split(authorization, " ")
		if authorization != "" && len(splited) == 2 {
			admin, err := adminService.VerifyToken(ctx, splited[1])
			if admin != nil && err == nil {
				ctx.Set("admin", admin)
			} else {
			}
		} else {
		}
		ctx.Next()
	}
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
