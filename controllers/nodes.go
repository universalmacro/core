package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/universalmacro/common/fault"
	"github.com/universalmacro/common/utils"
	"github.com/universalmacro/core/controllers/models"
	"github.com/universalmacro/core/dao/entities"
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

func (c *NodeController) GetNode(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	id := ctx.Param("id")
	node := c.NodeService.GetNode(utils.StringToUint(id))
	ctx.JSON(http.StatusOK, models.NodeConvertor(node))
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

func (c *NodeController) GetNodeConfig(ctx *gin.Context) {
	id := ctx.Param("id")
	node := c.NodeService.GetNode(utils.StringToUint(id))
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
	if admin.Role() != "ROOT" {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	ctx.JSON(http.StatusOK, models.NodeConfigConvertor(node.Config()))
}

func (c *NodeController) GetNodeDatabaseConfig(ctx *gin.Context) {
	id := ctx.Param("id")
	node := c.NodeService.GetNode(utils.StringToUint(id))
	if node == nil {
		fault.GinHandler(ctx, fault.ErrNotFound)
		return
	}
	dbConfig := node.GetDatabaseConfig()
	var headers Headers
	ctx.ShouldBindHeader(&headers)
	if headers.ApiKey != nil {
		if *headers.ApiKey == node.SecurityKey() {
			ctx.JSON(http.StatusOK, dbConfig)
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
	if admin.Role() != "ROOT" {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	if dbConfig == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}
	ctx.JSON(http.StatusOK, dbConfig)
}

func (c *NodeController) UpdateNodeDatabaseConfig(ctx *gin.Context) {
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
	node := c.NodeService.GetNode(utils.StringToUint(id))
	var dbConfig entities.DBConfig
	ctx.ShouldBindJSON(&dbConfig)
	node.UpdateDatabaseConfig(&dbConfig)
	ctx.JSON(http.StatusOK, dbConfig)
}

func (c *NodeController) GetNodeRedisConfig(ctx *gin.Context) {
	id := ctx.Param("id")
	node := c.NodeService.GetNode(utils.StringToUint(id))
	if node == nil {
		fault.GinHandler(ctx, fault.ErrNotFound)
		return
	}
	admin := getAdmin(ctx)
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return
	}
	if admin.Role() != "ROOT" {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	redisConfig := node.GetRedisConfig()
	if redisConfig == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}
	ctx.JSON(http.StatusOK, redisConfig)
}

func (c *NodeController) UpdateNodeRedisConfig(ctx *gin.Context) {
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
	node := c.NodeService.GetNode(utils.StringToUint(id))
	var redisConfig entities.RedisConfig
	ctx.ShouldBindJSON(&redisConfig)
	node.UpdateRedisConfig(&redisConfig)
	ctx.JSON(http.StatusOK, redisConfig)
}

func (c *NodeController) DeleteNode(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin == nil || admin.Role() != "ROOT" {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	id := ctx.Param("id")
	c.NodeService.DeleteNode(utils.StringToUint(id))
	ctx.JSON(http.StatusNoContent, nil)
}
