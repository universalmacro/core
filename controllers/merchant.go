package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/universalmacro/common/fault"
	"github.com/universalmacro/common/server"
	api "github.com/universalmacro/core-api-interfaces"
	"github.com/universalmacro/core/services"
)

func newMerchantController() *MerchantController {
	return &MerchantController{services.GetNodeService()}
}

type MerchantController struct {
	nodeService *services.NodeService
}

// DeleteMerchant implements coreapiinterfaces.MerchantApi.
func (*MerchantController) DeleteMerchant(ctx *gin.Context) {
	panic("unimplemented")
}

// GetMerchant implements coreapiinterfaces.MerchantApi.
func (*MerchantController) GetMerchant(ctx *gin.Context) {
	panic("unimplemented")
}

// AddMerchantToNode implements coreapiinterfaces.MerchantApi.
func (c *MerchantController) AddMerchantToNode(ctx *gin.Context) {
	admin := getAdmin(ctx)
	if admin == nil {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	var request api.CreateMerchantRequest
	ctx.ShouldBindJSON(&request)
	node := c.nodeService.GetNode(server.UintID(ctx, "id"))
	if node == nil {
		fault.GinHandler(ctx, fault.ErrNotFound)
		return
	}
	node.CreateMerchant(request.Account, request.Password)
	ctx.JSON(http.StatusCreated, api.Merchant{
		Id:   "1",
		Name: "Test Merchant",
	})
}

// ListNodeMerchants implements coreapiinterfaces.MerchantApi.
func (*MerchantController) ListNodeMerchants(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, api.MerchantList{
		Items: []api.Merchant{
			{
				Id:   "1",
				Name: "Test Merchant",
			},
			{
				Id:   "2",
				Name: "Test Merchant",
			},
		},
		Pagination: api.Pagination{
			Total: 2,
			Limit: 0,
			Index: 0,
		},
	})
}
