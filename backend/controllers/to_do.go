package controllers

import (
	"fmt"
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
}

func (c ToDoController) all(ctx *gin.Context) {
	fmt.Println("GET all")
	toDoModel := new(models.ToDoModel)
	ctx.JSON(http.StatusOK, toDoModel.GetAll())
}

func (c ToDoController) toggleComplete(ctx *gin.Context) {
	fmt.Println("POST markComplete")
	toDoModel := new(models.ToDoModel)
	err := toDoModel.ToggleComplete(ctx.Param("id"))

	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, toDoModel.GetAll())
	}
}
