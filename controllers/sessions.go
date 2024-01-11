package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/universalmacro/common/fault"
	"github.com/universalmacro/core/controllers/models"
	"github.com/universalmacro/core/services"
)

type SessionsController struct {
	adminService *services.AdminService
}

func newSessionsController() *SessionsController {
	return &SessionsController{adminService: services.GetAdminService()}
}

func (a *SessionsController) CreateSession(ctx *gin.Context) {
	var createSessionRequest models.CreateSessionRequest
	ctx.ShouldBindJSON(&createSessionRequest)
	token, err := a.adminService.CreateSession(createSessionRequest.Account, createSessionRequest.Password)
	if err != nil {
		fault.GinHandler(ctx, err)
		return
	}
	session := models.Session{Token: token}
	ctx.JSON(http.StatusCreated, session)
}
