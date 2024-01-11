package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/universalmacro/core/services"
)

type SessionsController struct {
	adminService *services.AdminService
}

func newSessionsController() *SessionsController {
	return &SessionsController{adminService: services.GetAdminService()}
}

func (a *SessionsController) CreateSession(ctx *gin.Context) {

}
