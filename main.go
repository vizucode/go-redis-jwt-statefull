package main

import (
	"go_jwt_statefull/config"
	"go_jwt_statefull/handlers"
	"go_jwt_statefull/middlewares"
	repositories "go_jwt_statefull/repositories"
	"go_jwt_statefull/services"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.NewDB()
	config.AutoMigrate(db)

	userRepo := repositories.NewUser(db)
	userService := services.NewUserImpl(userRepo)
	userHandler := handlers.NewUsers(userService)

	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.CustomRecovery(handlers.Recovery))

	userAPI := server.Group("/users", middlewares.Authorize)
	{
		userAPI.GET("/", userHandler.FindAll)
		userAPI.POST("/store", userHandler.Store)
	}

	server.NoRoute(handlers.NoRoute)

	port := os.Getenv("SERVER_PORT")
	if err := server.Run(":" + port); err != nil {
		panic(err.Error())
	}
}
