package main

import (
	"net/http"

	"github.com/Ricardo-DevNull/task-api/config"
	"github.com/Ricardo-DevNull/task-api/internal/infra/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Get server port
	port := ":" + config.SvConfig.Port

	// Start Engine Gin Framework
	router := gin.Default()

	// Init Database
	database.InitDB()

	// Execute migrations
	database.Migrate()

	router.GET("/task-api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	if err := router.Run(port); err != nil {
		return
	}
}
