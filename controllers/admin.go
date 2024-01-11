package controllers

import (
	"github.com/universalmacro/core/services"
)

type AdminController struct {
	adminService *services.AdminService
}

func newAdminController() *AdminController {
	return &AdminController{adminService: services.GetAdminService()}
}
