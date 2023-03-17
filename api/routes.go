package api

import (
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/controller"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/repository"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GinRoute(db *gorm.DB) *gin.Engine {

	r := gin.New()
	r.HandleMethodNotAllowed = true

	userRepository := repository.NewUserRepositoryImplementation()
	userService := service.NewUserServiceImplementation(db, userRepository)
	userController := controller.NewUserControllerImplementation(userService)

	authRepository := repository.NewAuthRepositoryImplementation()
	authService := service.NewAuthServiceImplementation(db, authRepository, userService)
	authController := controller.NewAuthControllerImplementation(authService)

	categoryRepository := repository.NewCategoryRepositoryImplementation()
	categoryService := service.NewCategoryServiceImplementation(db, categoryRepository)
	categoryController := controller.NewCategoryControllerImplementation(categoryService)

	api := r.Group("api/v.1")
	user := api.Group("/user")
	user.POST("/add", userController.AddUser)
	user.PUT("/update", userController.UpdateUser)
	user.DELETE("/delete", userController.DeleteUser)
	user.GET("/get", func(c *gin.Context) {
		if c.Query("id") != "" {
			userController.GetUserByID(c)
			return
		} else if c.Query("npm") != "" {
			userController.GetUserByNPM(c)
			return
		}
	})

	auth := api.Group("/auth")
	auth.POST("/register", authController.AuthRegister)
	auth.POST("/login", authController.AuthLogin)

	category := api.Group("/category")
	category.POST("/add", categoryController.AddCategory)
	category.PUT("/update", categoryController.UpdateCategory)
	category.DELETE("/delete", categoryController.DeleteCategory)
	category.GET("/get", func(c *gin.Context) {
		if c.Query("id") != "" {
			categoryController.GetCategoryByID(c)
			return
		} else if c.Query("name") != "" {
			categoryController.GetCategoryByName(c)
			return
		} else {
			categoryController.GetCategories(c)
			return
		}

	})

	return r
}
