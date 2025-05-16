package main

import (
	"demo/internal/api"
	"demo/internal/configs"
	"demo/internal/repositories"
	"demo/internal/services/links"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := configs.ConnectDB()
	defer db.Close()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	linkRepo := repositories.NewLinkRepository(db)
	linkService := links.NewLinkService(linkRepo)
	linkAPI := api.NewLinkHandler(linkService)

	linkGroup := router.Group("/links")
	{
		linkGroup.GET("", linkAPI.GetAll)
		linkGroup.POST("/add", linkAPI.CreateLink)
		linkGroup.GET("/:id", linkAPI.GetByID)
	}

	router.GET("/:code", linkAPI.GetByCode)

	router.Run("0.0.0.0:1234")
}
