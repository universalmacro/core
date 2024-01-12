package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/universalmacro/common/fault"
	"github.com/universalmacro/core/controllers/models"
	"github.com/universalmacro/core/services"
)

func newAdminController() *AdminController {
	return &AdminController{adminService: services.GetAdminService()}
}

type AdminController struct {
	adminService *services.AdminService
}

func (a *AdminController) CreateAdmin(ctx *gin.Context) {
	account := getAccount(ctx)
	if account.Role() != "ROOT" {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	var createAdminRequest models.CreateAdminRequest
	ctx.ShouldBindJSON(&createAdminRequest)
	admin, err := a.adminService.CreateAdmin(createAdminRequest.Account, createAdminRequest.Password, createAdminRequest.Role)
	if err != nil {
		fault.GinHandler(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, models.AdminConvertor(*admin))
}
