package main

import (
	"github.com/Ricardo-DevNull/task-api/config"
	"github.com/Ricardo-DevNull/task-api/internal/infra/database"
	"github.com/Ricardo-DevNull/task-api/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// Get server port
	port := ":" + config.SvConfig.Port

	// Start Engine Gin Framework
	r := gin.Default()

	// Init Database
	db := database.InitDB()

	// Execute migrations
	database.Migrate(db)

	// router.SetupRoutes(router, db)
	router.SetupRoutes(r, db)

	if err := r.Run(port); err != nil {
		return
	}
}
