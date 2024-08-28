package controllers

import (
	"io"
	"net/http"
	"oneblock_manage_fund/security"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
}

func (o *LoginController) Router(r *gin.RouterGroup) {
	r.GET("/login", o.Login)
}
func (o *LoginController) Login(c *gin.Context) {
	tokenResp := security.TokenResponse{}
	data, _ := io.ReadAll(c.Request.Body)
	tokenResp.Token = string(data)
	c.JSON(http.StatusOK, tokenResp)
}
