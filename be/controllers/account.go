package controllers

import "github.com/gin-gonic/gin"

type Account struct{}

func (a *Account) Router(r *gin.RouterGroup) {
	r.GET("/login", a.Authenticate)
}

func (a *Account) Authenticate(c *gin.Context) {}
