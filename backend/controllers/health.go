package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthController ...
type HealthController struct{}

// RegisterRoutes ...
func (c HealthController) RegisterRoutes(r *gin.Engine) {
	r.GET("/health-check", c.healthCheck)
}

func (c HealthController) healthCheck(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Healthy")
}
