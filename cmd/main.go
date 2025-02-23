package main

import (
	"mtuan8820/go-todo-list/v2/controller"
	"mtuan8820/go-todo-list/v2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	taskService    service.TaskService       = service.New()
	taskController controller.TaskController = controller.New(taskService)
)

func main() {
	server := gin.Default()
	server.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(200, taskController.GetAll())
	})

	server.POST("/tasks", func(ctx *gin.Context) {
		err := taskController.Add(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Task created successfully"})
		}
	})

	server.DELETE("/tasks/:id", func(ctx *gin.Context) {
		err := taskController.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
		}
	})

	server.PUT("/tasks/:id", func(ctx *gin.Context) {
		err := taskController.Toggle(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Task toggled successfully"})
		}
	})

	server.Run(":8080")
}
