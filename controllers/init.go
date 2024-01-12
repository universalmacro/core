package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/universalmacro/common/fault"
	"github.com/universalmacro/common/server"
	"github.com/universalmacro/core/services"
	"github.com/universalmacro/core/services/models"
)

var router = gin.Default()

var version = "0.0.2"

type Headers struct {
	Authorization string
}

func Init(addr ...string) {
	var adminController = newAdminController()
	var adminService = services.GetAdminService()
	var sessionsControllers = newSessionsController()
	router.Use(server.CorsMiddleware())
	router.Use(func(ctx *gin.Context) {
		var headers Headers
		ctx.ShouldBindHeader(&headers)
		authorization := headers.Authorization
		splited := strings.Split(authorization, " ")
		if authorization != "" && len(splited) == 2 {
			admin, _ := adminService.VerifyToken(splited[1])
			if admin != nil {
				ctx.Set("admin", admin)
			}
		}
		ctx.Next()
	})
	server.MetricsMiddleware(router)
	router.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version": version,
		})
	})
	// Session
	router.POST("/sessions", sessionsControllers.CreateSession)
	router.GET("/test", func(ctx *gin.Context) {
		admin := getAccount(ctx)
		fmt.Println(admin)
	})
	// Admin
	router.POST("/admins", adminController.CreateAdmin)
	router.Run(addr...)
}

func getAccount(ctx *gin.Context) *models.Admin {
	adminInterface, ok := ctx.Get("account")
	if !ok {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return nil
	}
	admin, ok := adminInterface.(models.Admin)
	if !ok {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return nil
	}
	return &admin
}
