package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/universalmacro/common/fault"
	"github.com/universalmacro/common/server"
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

// GetNodeApiConfigByDomain implements coreapiinterfaces.NodeApi.
func (c *NodeController) GetNodeApiConfigByDomain(ctx *gin.Context) {
	domain := ctx.Query("domain")
	node := c.NodeService.GetNodeByFrontendDomain(domain)
	if node == nil {
		fault.GinHandler(ctx, fault.ErrNotFound)
		return
	}
	ctx.JSON(http.StatusOK, node.Config().Api)
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

func (c *NodeController) GetNode(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	node := c.NodeService.GetNode(server.UintID(ctx, "id"))
	ctx.JSON(http.StatusOK, models.NodeConvertor(node))
}

func (c *NodeController) ListNode(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	index, limit := server.IndexAndLimit(ctx)
	nodeList := c.NodeService.ListNode(index, limit)
	ctx.JSON(http.StatusOK, models.NodeListConvertor(nodeList))
}

func (c *NodeController) GetNodeConfig(ctx *gin.Context) {
	node := c.NodeService.GetNode(server.UintID(ctx, "id"))
	if node == nil {
		fault.GinHandler(ctx, fault.ErrNotFound)
		return
	}
	var headers Headers
	ctx.ShouldBindHeader(&headers)
	if headers.ApiKey != nil {
		if *headers.ApiKey == node.SecurityKey() {
			ctx.JSON(http.StatusOK, models.NodeConfigConvertor(node.Config()))
		} else {
			fault.GinHandler(ctx, fault.ErrUnauthorized)
		}
		return
	}
	admin := getAdmin(ctx)
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return
	}
	ctx.JSON(http.StatusOK, models.NodeConfigConvertor(node.Config()))
}

func (c *NodeController) UpdateNodeConfig(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return
	}
	node := c.NodeService.GetNode(server.UintID(ctx, "id"))
	if node == nil {
		fault.GinHandler(ctx, fault.ErrNotFound)
		return
	}
	var nodeConfig models.NodeConfig
	ctx.ShouldBindJSON(&nodeConfig)
	node.UpdateConfig(nodeConfig.FrontendDomain, nodeConfig.Api, nodeConfig.Server, nodeConfig.Database, nodeConfig.Redis, nodeConfig.TencentCloud)
	ctx.JSON(http.StatusOK, models.NodeConfigConvertor(node.Config()))
}

func (c *NodeController) DeleteNode(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin == nil || admin.Role() != "ROOT" {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	c.NodeService.DeleteNode(server.UintID(ctx, "id"))
	ctx.JSON(http.StatusNoContent, nil)
}
