package controllers

import "github.com/gin-gonic/gin"

// HomeController ...
type HomeController struct{}

// RegisterRoutes ...
func (c HomeController) RegisterRoutes(r *gin.Engine) {
	r.GET("/", c.Index)
}

// Index ...
func (c HomeController) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"welcome": "Home",
	})
}
