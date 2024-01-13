package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/universalmacro/common/fault"
	"github.com/universalmacro/core/controllers/models"
	"github.com/universalmacro/core/services"
)

func newNodeController() *NodeController {
	return &NodeController{
		NodeService: services.GetNodeService(),
	}
}

type NodeController struct {
	NodeService *services.NodeService
}

func (c *NodeController) CreateNode(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin.Role() != "ROOT" {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	var createNodeRequest models.CreateNodeRequest
	ctx.ShouldBindJSON(&createNodeRequest)
}
