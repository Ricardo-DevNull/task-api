package router

import (
	"github.com/Ricardo-DevNull/task-api/internal/handlers"
	"github.com/Ricardo-DevNull/task-api/internal/repository"
	"github.com/Ricardo-DevNull/task-api/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api/v1")

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	userRoutes := api.Group("/users")
	{
		userRoutes.POST("", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	taskRepo := repository.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskService(taskService)

	taskRouter := api.Group("/tasks")
	{
		taskRouter.POST("", taskHandler.CreateTask)
	}
}
