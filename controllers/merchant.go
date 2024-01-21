package controllers

import "github.com/gin-gonic/gin"

func newMerchantController() *MerchantController {
	return &MerchantController{}
}

type MerchantController struct{}

// AddMerchantToNode implements coreapiinterfaces.MerchantApi.
func (*MerchantController) AddMerchantToNode(ctx *gin.Context) {
	panic("unimplemented")
}

// ListNodeMerchants implements coreapiinterfaces.MerchantApi.
func (*MerchantController) ListNodeMerchants(ctx *gin.Context) {
	panic("unimplemented")
}
