package controller

import (
	"mtuan8820/go-todo-list/v2/pkg/models"
	"mtuan8820/go-todo-list/v2/service"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	FindAll() []models.Task
	Save(ctx *gin.Context) models.Task
}

type controller struct {
	service service.TaskService
}

func New(service service.TaskService) TaskController {
	return &controller{service: service}
}

func (c *controller) FindAll() []models.Task {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) models.Task {
	var task models.Task
	ctx.BindJSON(&task)
	c.service.Save(task)
	return task
}
