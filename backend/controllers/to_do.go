package controllers

import (
	"net/http"
	"todone-backend/models"

	"github.com/gin-gonic/gin"
)

// ToDoController ...
type ToDoController struct{}

// RegisterRoutes ...
func (c ToDoController) RegisterRoutes(r *gin.Engine) {
	r.GET("/todo", c.all)
}

func (c ToDoController) all(ctx *gin.Context) {
	toDoModel := new(models.ToDoModel)
	ctx.JSON(http.StatusOK, toDoModel.GetAll())
}
