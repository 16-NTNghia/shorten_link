package main

import (
	"demo/internal/api"
	"demo/internal/configs"
	"demo/internal/mappers"
	"demo/internal/repositories"
	"demo/internal/services"
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
	linkService := services.NewLinkService(linkRepo)
	linkAPI := api.NewLinkHandler(linkService)

	userRepo := repositories.NewUserRepository(db)
	userMapper := mappers.NewUserMapper()
	userService := services.NewUserService(userRepo, userMapper)
	userAPI := api.NewUsersHandler(userService)

	linkGroup := router.Group("/links")
	{
		linkGroup.GET("", linkAPI.GetAll)
		linkGroup.POST("/add", linkAPI.CreateLink)
		linkGroup.GET("/:id", linkAPI.GetByID)
	}

	userGroup := router.Group("/users")
	{
		userGroup.GET("", userAPI.GetAll)
		userGroup.POST("/add", userAPI.CreateNewUser)
		userGroup.GET("/:id", userAPI.GetByID)
		userGroup.PUT("/:id", userAPI.UpdateUser)
	}

	router.GET("/:code", linkAPI.GetByCode)

	router.Run("0.0.0.0:1234")
}
