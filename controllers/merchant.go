package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/universalmacro/core-api-interfaces"
)

func newMerchantController() *MerchantController {
	return &MerchantController{}
}

type MerchantController struct{}

// AddMerchantToNode implements coreapiinterfaces.MerchantApi.
func (*MerchantController) AddMerchantToNode(ctx *gin.Context) {
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
