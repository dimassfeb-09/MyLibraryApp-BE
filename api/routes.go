package api

import (
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/controller"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/repository"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GinRoute(db *gorm.DB) *gin.Engine {

	r := gin.Default()
	r.HandleMethodNotAllowed = true

	userRepository := repository.NewUserRepositoryImplementation()
	authRepository := repository.NewAuthRepositoryImplementation()
	categoryRepository := repository.NewCategoryRepositoryImplementation()
	bookRepository := repository.NewBookRepositoryImplementation()
	wishlistRepository := repository.NewWishlistRepositoryImplementation()
	ratingRepository := repository.NewRatingRepositoryImplementation()
	genreRepository := repository.NewGenreRepositoryImplementation()

	microService := repository.NewRegisterMicroServiceImplementation(userRepository, authRepository, categoryRepository, bookRepository, wishlistRepository, ratingRepository, genreRepository)

	userService := service.NewUserServiceImplementation(db, microService)
	authService := service.NewAuthServiceImplementation(db, microService)
	categoryService := service.NewCategoryServiceImplementation(db, microService)
	bookService := service.NewBookServiceImplementation(db, microService)
	wishlistService := service.NewWishlistServiceImplementation(db, microService)
	ratingService := service.NewRatingServiceImplementation(db, microService)
	genreService := service.NewGenreServiceImplementation(db, microService)

	userController := controller.NewUserControllerImplementation(userService)
	authController := controller.NewAuthControllerImplementation(authService)
	categoryController := controller.NewCategoryControllerImplementation(categoryService)
	bookController := controller.NewBookControllerImplementation(bookService)
	wishlistController := controller.NewWishlistControllerImplementation(wishlistService)
	ratingController := controller.NewRatingControllerImplementation(ratingService)
	genreController := controller.NewGenreControllerImplementation(genreService)

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

	book := api.Group("/book")
	book.POST("/add", bookController.AddBook)
	book.PUT("/update", bookController.UpdateBook)
	book.DELETE("/delete", bookController.DeleteBook)
	book.GET("/get", func(c *gin.Context) {
		if c.Query("id") != "" {
			bookController.GetBookByID(c)
			return
		} else if c.Query("title") != "" {
			bookController.GetBookByTitle(c)
			return
		} else if c.Query("genre_id") != "" {
			bookController.GetBooksByGenreID(c)
			return
		} else if c.Query(("category_id")) != "" {
			bookController.GetBooksByCategoryID(c)
			return
		} else {
			bookController.GetBooks(c)
			return
		}

	})

	wishlist := api.Group("/wishlist")
	wishlist.POST("/add", wishlistController.AddWishlist)
	wishlist.PUT("/update", wishlistController.UpdateWishlist)
	wishlist.DELETE("/delete", wishlistController.DeleteWishlist)
	wishlist.GET("/get", func(c *gin.Context) {
		if c.Query("id") != "" {
			wishlistController.GetWishlistByID(c)
			return
		} else if c.Query("user_id") != "" {
			wishlistController.GetWishlistByUserID(c)
			return
		} else {
			wishlistController.GetWishlists(c)
			return
		}
	})

	rating := api.Group("/rating")
	rating.POST("/add", ratingController.AddRating)
	rating.PUT("/update", ratingController.UpdateRating)
	rating.DELETE("/delete", ratingController.DeleteRating)
	rating.GET("/get", func(c *gin.Context) {
		if c.Query("id") != "" {
			ratingController.GetRatingByID(c)
			return
		} else if c.Query("book_id") != "" {
			ratingController.GetRatingByBookID(c)
			return
		}
	})

	genre := api.Group("/genre")
	genre.POST("/add", genreController.AddGenre)
	genre.PUT("/update", genreController.UpdateGenre)
	genre.DELETE("/delete", genreController.DeleteGenre)
	genre.GET("/get", func(c *gin.Context) {
		if c.Query("id") != "" {
			genreController.GetGenreByID(c)
			return
		} else if c.Query(("category_id")) != "" {
			genreController.GetGenreByCategoryID(c)
			return
		} else {
			genreController.GetGenres(c)
			return
		}
	})

	return r
}
