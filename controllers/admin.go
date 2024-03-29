package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
	"github.com/universalmacro/common/fault"
	"github.com/universalmacro/common/utils"
	api "github.com/universalmacro/core-api-interfaces"
	"github.com/universalmacro/core/controllers/models"
	"github.com/universalmacro/core/services"
)

func newAdminController() *AdminController {
	return &AdminController{adminService: services.GetAdminService()}
}

type AdminController struct {
	adminService *services.AdminService
}

// CreateTotp implements coreapiinterfaces.AdminApi.
func (*AdminController) UpdateTotp(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return
	}
	var request api.UpdateTotpRequest
	ctx.ShouldBindJSON(&request)
	if ok := admin.UpdateTotp(request.Url, request.TotpCode); !ok {
		fault.GinHandler(ctx, fault.ErrBadRequest)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// GetTotp implements coreapiinterfaces.AdminApi.
func (*AdminController) GetTotp(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return
	}
	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "universalmacro.com",
		AccountName: admin.Account(),
	})
	ctx.JSON(http.StatusOK, api.Totp{
		Url: key.URL(),
	})
}

// UpdateAdminRole implements coreapiinterfaces.AdminApi.
func (*AdminController) UpdateAdminRole(ctx *gin.Context) {
	panic("unimplemented")
}

// GetAdminSelf implements coreapiinterfaces.AdminApi.
func (*AdminController) GetAdminSelf(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return
	}
	ctx.JSON(http.StatusOK, models.AdminConvertor(*admin))
}

// UpdateAdminPassword implements coreapiinterfaces.AdminApi.
func (c *AdminController) UpdateAdminPassword(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return
	}
	if admin.Role() != "ROOT" {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	id := ctx.Param("id")
	var updatePasswordRequest models.UpdatePasswordRequest
	ctx.ShouldBindJSON(&updatePasswordRequest)
	account := c.adminService.UpdatePassword(utils.StringToUint(id), updatePasswordRequest.Password)
	ctx.JSON(http.StatusOK, models.AdminConvertor(*account))
}

// UpdateAdminSelfPassword implements coreapiinterfaces.AdminApi.
func (c *AdminController) UpdateAdminSelfPassword(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return
	}
	var updatePasswordRequest models.UpdatePasswordRequest
	ctx.ShouldBindJSON(&updatePasswordRequest)
	if updatePasswordRequest.OldPassword == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": services.ErrPasswordNotMatch,
		})
		return
	}
	err := c.adminService.UpdateSelfPassword(ctx, utils.UintToString(admin.ID()), *updatePasswordRequest.OldPassword, updatePasswordRequest.Password)
	if err != nil {
		fault.GinHandler(ctx, err)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (a *AdminController) CreateAdmin(ctx *gin.Context) {
	account := getAdmin(ctx)
	if account == nil || account.Role() != "ROOT" {
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

func (a *AdminController) GetAdmin(ctx *gin.Context) {
	admin := getAdmin(ctx)
	id := ctx.Param("id")
	a.adminService.GetAdmin(utils.StringToUint(id))
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return
	}
	ctx.JSON(http.StatusOK, models.AdminConvertor(*admin))
}

func (c *AdminController) ListAdmin(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	index := ctx.Query("index")
	limit := ctx.Query("limit")
	adminList := c.adminService.ListAdmin(int64(utils.StringToUint(index)), int64(utils.StringToUint(limit)))
	ctx.JSON(http.StatusOK, models.AdminListConvertor(adminList))
}

func (c *AdminController) DeleteAdmin(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return
	}
	if admin.Role() != "ROOT" {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	id := ctx.Param("id")
	err := c.adminService.DeleteAdmin(utils.StringToUint(id))
	if err != nil {
		fault.GinHandler(ctx, err)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
