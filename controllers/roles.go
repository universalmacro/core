package controllers

import "github.com/gin-gonic/gin"

func newRolesController() *RolesController {
	return &RolesController{}
}

type RolesController struct {
}

// CreateRole implements coreapiinterfaces.RoleApi.
func (*RolesController) CreateRole(ctx *gin.Context) {
	panic("unimplemented")
}

// DeleteRole implements coreapiinterfaces.RoleApi.
func (*RolesController) DeleteRole(ctx *gin.Context) {
	panic("unimplemented")
}

// GetRole implements coreapiinterfaces.RoleApi.
func (*RolesController) GetRole(ctx *gin.Context) {
	panic("unimplemented")
}

// GetRolePermissions implements coreapiinterfaces.RoleApi.
func (*RolesController) GetRolePermissions(ctx *gin.Context) {
	panic("unimplemented")
}

// ListRoles implements coreapiinterfaces.RoleApi.
func (*RolesController) ListRoles(ctx *gin.Context) {
	panic("unimplemented")
}

// UpdateRole implements coreapiinterfaces.RoleApi.
func (*RolesController) UpdateRole(ctx *gin.Context) {
	panic("unimplemented")
}

// UpdateRolePermissions implements coreapiinterfaces.RoleApi.
func (*RolesController) UpdateRolePermissions(ctx *gin.Context) {
	panic("unimplemented")
}
