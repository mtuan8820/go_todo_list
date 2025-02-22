package main

import (
	"mtuan8820/go-todo-list/v2/controller"
	"mtuan8820/go-todo-list/v2/service"

	"github.com/gin-gonic/gin"
)

var (
	taskService    service.TaskService       = service.New()
	taskController controller.TaskController = controller.New(taskService)
)

func main() {
	server := gin.Default()
	server.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(200, taskController.FindAll())
	})

	server.POST("/tasks", func(ctx *gin.Context) {
		ctx.JSON(200, taskController.Save(ctx))
	})

	server.Run(":8080")
}
