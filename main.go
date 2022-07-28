package main

import (
	"go_jwt_statefull/config"
	"go_jwt_statefull/handlers"
	"go_jwt_statefull/middlewares"
	repositories "go_jwt_statefull/repositories"
	"go_jwt_statefull/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.NewDB()
	userRepo := repositories.NewUser(db)
	authService := services.NewAuthentication(userRepo)
	authHandler := handlers.NewAuth(authService)

	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.CustomRecovery(handlers.Recovery))

	server.POST("/login", authHandler.Login)

	logout := server.Group("/logout", middlewares.Authorization)
	{
		logout.POST("/", authHandler.Logout)
	}

	server.NoRoute(handlers.NoRoute)

	if err := server.Run(":8080"); err != nil {
		panic(err.Error())
	}
}
