package controllers

import (
	"net/http"
	"oneblock_manage_fund/models"
	"oneblock_manage_fund/thirdparty"

	"github.com/gin-gonic/gin"
)

type BinanceController struct{}

func (a *BinanceController) Router(r *gin.RouterGroup) {
	r.GET("/balance", a.WalletBalance)
	r.GET("/asset", a.WalletAsset)
}

func (a *BinanceController) WalletBalance(ctx *gin.Context) {
	balance, err := thirdparty.UserBalance()
	if err != nil {
		balance = make(models.BinanceWalletBalance, 0)
	}
	ctx.JSON(http.StatusOK, balance)
}

func (a *BinanceController) WalletAsset(ctx *gin.Context) {
	balance, err := thirdparty.UserAsset()
	if err != nil {
		balance = make(models.BinanceUserAsset, 0)
	}
	ctx.JSON(http.StatusOK, balance)
}
