package main

import (
	"go_jwt_statefull/config"
	"go_jwt_statefull/handlers"
	repositories "go_jwt_statefull/repositories"
	"go_jwt_statefull/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.NewDB()
	userRepo := repositories.NewUser(db)
	userService := services.NewAuthentication(userRepo)
	userHandler := handlers.NewAuth(userService)

	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.CustomRecovery(handlers.Recovery))

	server.POST("/login", userHandler.Login)

	server.NoRoute(handlers.NoRoute)

	if err := server.Run(":8080"); err != nil {
		panic(err.Error())
	}
}
