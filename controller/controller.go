package controller

import (
	"mtuan8820/go-todo-list/v2/pkg/models"
	"mtuan8820/go-todo-list/v2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	GetAll() []models.Task
	Add(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	Toggle(ctx *gin.Context) error
	Show(ctx *gin.Context)
}

type controller struct {
	service service.TaskService
}

func New(service service.TaskService) TaskController {
	return &controller{service: service}
}

func (c *controller) GetAll() []models.Task {
	return c.service.GetAll()
}

func (c *controller) Add(ctx *gin.Context) error {
	var task models.Task
	err := ctx.BindJSON(&task)
	if err != nil {
		return (err)
	}
	c.service.Add(task)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	id := ctx.Param("id")
	// convert string id to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = c.service.Delete(intID)
	if err != nil {
		return (err)
	}
	return nil
}

func (c *controller) Toggle(ctx *gin.Context) error {
	id := ctx.Param("id")
	// convert string id to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = c.service.Toggle(intID)
	if err != nil {
		return (err)
	}
	return nil
}

func (c *controller) Show(ctx *gin.Context) {
	tasks := c.service.GetAll()
	data := gin.H{
		"title": "Todo List",
		"tasks": tasks,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}
