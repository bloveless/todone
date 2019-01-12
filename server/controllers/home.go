package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeController ...
type HomeController struct{}

// RegisterRoutes ...
func (c HomeController) RegisterRoutes(r *gin.Engine) {
	r.GET("/", c.index)
}

// Index ...
func (c HomeController) index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"welcome": "Home",
	})
}
