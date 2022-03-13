// routes/routes.go
package routes

import (
	"taskcrud/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"taskcrud/middlewares"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	v1 := r.Group("/v1")
	{

		v1.POST("/users/register", controllers.CreateUser)
		v1.POST("/users/login", controllers.Login)
		v1.POST("/users/reset-password", controllers.ResetPassword)
		v1.GET("/users", middlewares.JwtAuthMiddleware(), controllers.FindUsers)
		v1.GET("/users/:id", middlewares.JwtAuthMiddleware(), controllers.FindUser)
		v1.PATCH("/users/:id", middlewares.JwtAuthMiddleware(), controllers.UpdateUser)
		v1.DELETE("users/:id", middlewares.JwtAuthMiddleware(), controllers.DeleteUser)
		v1.POST("users/verification/*token", controllers.VerificationUser)

		v1.GET("/tasks", middlewares.JwtAuthMiddleware(), controllers.FindTasks)
		v1.POST("/tasks", middlewares.JwtAuthMiddleware(), controllers.CreateTask)
		v1.GET("/tasks/:id", middlewares.JwtAuthMiddleware(), controllers.FindTask)
		v1.PATCH("/tasks/:id", middlewares.JwtAuthMiddleware(), middlewares.TaskOwnershipValidation(), controllers.UpdateTask)
		v1.DELETE("tasks/:id", middlewares.JwtAuthMiddleware(), middlewares.TaskOwnershipValidation(), controllers.DeleteTask)

	}

	return r
}
