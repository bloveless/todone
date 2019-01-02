package controllers

import (
	"github.com/gin-gonic/gin"
)

// Pong returns a message
type Pong struct {
	Message string `json:"message"`
}

// PingController ...
type PingController struct{}

// RegisterRoutes ...
func (c PingController) RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", c.Pong)
}

// Pong ...
func (c PingController) Pong(ctx *gin.Context) {
	pong := Pong{"Pong"}
	ctx.JSON(200, pong)
}
