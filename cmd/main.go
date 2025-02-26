package main

import (
	"log"
	"mtuan8820/go-todo-list/v2/controller"
	"mtuan8820/go-todo-list/v2/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	taskService    service.TaskService       = service.New()
	taskController controller.TaskController = controller.New(taskService)
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.Use(CORSMiddleware())
	server.Static("/static", "../web/static")
	server.LoadHTMLGlob("../web/templates/*.html")

	apiRoutes := server.Group("/")
	{
		apiRoutes.GET("/tasks", func(ctx *gin.Context) {
			ctx.JSON(200, taskController.GetAll())
		})

		apiRoutes.POST("/tasks", func(ctx *gin.Context) {
			err := taskController.Add(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Task created successfully"})
			}
		})

		apiRoutes.DELETE("/tasks/:id", func(ctx *gin.Context) {
			err := taskController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
			}
		})

		apiRoutes.PUT("/tasks/:id", func(ctx *gin.Context) {
			err := taskController.Toggle(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Task toggled successfully"})
			}
		})

	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/tasks", taskController.Show)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := server.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
