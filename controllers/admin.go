package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/universalmacro/common/fault"
	"github.com/universalmacro/common/utils"
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
	account := getAdmin(ctx)
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

func (a *AdminController) GetSelf(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return
	}
	ctx.JSON(http.StatusOK, models.AdminConvertor(*admin))
}

func (c *AdminController) ListAdmin(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin.Role() != "ROOT" {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	index := ctx.Query("index")
	limit := ctx.Query("limit")
	adminList := c.adminService.ListAdmin(int64(utils.StringToUint(index)), int64(utils.StringToUint(limit)))
	ctx.JSON(http.StatusOK, models.AdminListConvertor(adminList))
}
