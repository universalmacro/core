package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/universalmacro/common/fault"
	"github.com/universalmacro/common/utils"
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
	node := c.NodeService.CreateNode(createNodeRequest.Name, createNodeRequest.Description)
	ctx.JSON(http.StatusCreated, models.NodeConvertor(node))
}

func (c *NodeController) ListNode(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	index := ctx.Query("index")
	limit := ctx.Query("limit")
	nodeList := c.NodeService.ListNode(int64(utils.StringToUint(index)), int64(utils.StringToUint(limit)))
	ctx.JSON(http.StatusOK, models.NodeListConvertor(nodeList))
}
