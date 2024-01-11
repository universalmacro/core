package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/universalmacro/core/services"
)

type AdminController struct {
	adminService *services.AdminService
}

func newAdminController() *AdminController {
	return &AdminController{adminService: services.GetAdminService()}
}

func (a *AdminController) Login(ctx *gin.Context) {

}
