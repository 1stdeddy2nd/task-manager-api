package main

import (
	"task-manager/controllers"
	"task-manager/database"
	"task-manager/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.POST("/api/auth/register", controllers.Register)
	r.POST("/api/auth/login", controllers.Login)

	auth := r.Group("/api/tasks")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/", controllers.GetTasks)
		auth.GET("/:id", controllers.GetTaskByID)
		auth.POST("/", controllers.CreateTask)
		auth.PUT("/:id", controllers.UpdateTask)
		auth.DELETE("/:id", controllers.DeleteTask)
	}

	r.Run(":8080")
}
