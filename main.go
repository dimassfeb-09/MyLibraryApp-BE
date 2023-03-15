package main

import (
	"fmt"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/api"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/controller"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/repository"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/service"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")
	db := api.DBConn()

	userRepository := repository.NewUserRepositoryImplementation()
	userService := service.NewUserServiceImplementation(db, userRepository)
	userController := controller.NewUserControllerImplementation(userService)

	r := gin.New()
	auth := r.Group("api/v.1/auth")
	auth.POST("/add", userController.AddUser)
	auth.GET("/get", userController.GetUserByID)

	r.Run(":8080")
}
