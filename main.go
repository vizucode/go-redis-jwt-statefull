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
	config.AutoMigrate(db)

	userRepo := repositories.NewUser(db)
	userService := services.NewUserImpl(userRepo)
	userHandler := handlers.NewUsers(userService)

	server := gin.New()
	server.Use(gin.Logger())

	userAPI := server.Group("/users")
	{
		userAPI.GET("/", userHandler.FindAll)
		userAPI.POST("/store", userHandler.Store)
	}

	server.NoRoute(handlers.NoRoute)

	if err := server.Run(":8080"); err != nil {
		panic(err.Error())
	}
}
