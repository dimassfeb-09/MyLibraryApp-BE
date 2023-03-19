package controller

import (
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/request"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/helpers"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookController interface {
	AddBook(c *gin.Context)
	UpdateBook(c *gin.Context)
	DeleteBook(c *gin.Context)
	GetBookByID(c *gin.Context)
	GetBookByTitle(c *gin.Context)
	GetBooks(c *gin.Context)
}

type BookControllerImplementation struct {
	BookService service.BookService
}

func NewBookControllerImplementation(bookService service.BookService) BookController {
	return &BookControllerImplementation{BookService: bookService}
}

func (b *BookControllerImplementation) AddBook(c *gin.Context) {
	var book *request.Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, err.Error(), nil))
		return
	}

	isSuccess, msg, err := b.BookService.AddBook(c.Request.Context(), book)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (b *BookControllerImplementation) UpdateBook(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	var book *request.Book
	err = c.ShouldBind(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, err.Error(), nil))
		return
	}
	book.ID = ID

	isSuccess, msg, err := b.BookService.UpdateBook(c.Request.Context(), book)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (b *BookControllerImplementation) DeleteBook(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	isSuccess, msg, err := b.BookService.DeleteBook(c.Request.Context(), ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}
	if isSuccess {
		c.JSON(http.StatusOK, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (b *BookControllerImplementation) GetBookByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	result, msg, err := b.BookService.GetBookByID(c.Request.Context(), ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}
	if result != nil {
		c.JSON(http.StatusOK, helpers.ToWebResponse("OK", http.StatusOK, msg, result))
		return
	}
}

func (b *BookControllerImplementation) GetBookByTitle(c *gin.Context) {
	title := c.Query("title")

	result, msg, err := b.BookService.GetBookByTitle(c.Request.Context(), title)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	c.JSON(http.StatusOK, helpers.ToWebResponse("OK", http.StatusOK, msg, result))
	return
}

func (b *BookControllerImplementation) GetBooks(c *gin.Context) {
	results, msg, err := b.BookService.GetBooks(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if results != nil {
		c.JSON(http.StatusOK, helpers.ToWebResponse("OK", http.StatusOK, msg, results))
		return
	}
}