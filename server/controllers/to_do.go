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
	r.GET("/to-do", c.all)
	r.POST("/to-do/:id/toggle-complete", c.toggleComplete)
	r.POST("/to-do/:id/toggle-active", c.toggleActive)
}

func (c ToDoController) all(ctx *gin.Context) {
	toDoModel := new(models.ToDoModel)
	ctx.JSON(http.StatusOK, toDoModel.GetAll())
}

func (c ToDoController) toggleComplete(ctx *gin.Context) {
	toDoModel := new(models.ToDoModel)
	err := toDoModel.ToggleComplete(ctx.Param("id"))

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
	} else {
		ctx.JSON(http.StatusOK, toDoModel.GetAll())
	}
}

func (c ToDoController) toggleActive(ctx *gin.Context) {
	toDoModel := new(models.ToDoModel)
	err := toDoModel.ToggleActive(ctx.Param("id"))

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
	} else {
		ctx.JSON(http.StatusOK, toDoModel.GetAll())
	}
}
