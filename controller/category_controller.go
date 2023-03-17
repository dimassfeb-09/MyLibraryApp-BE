package controller

import (
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/request"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/helpers"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryController interface {
	AddCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	GetCategoryByID(c *gin.Context)
	GetCategoryByName(c *gin.Context)
	GetCategories(c *gin.Context)
}

type CategoryControllerImplementation struct {
	CategoryService service.CategoryService
}

func NewCategoryControllerImplementation(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImplementation{CategoryService: categoryService}
}

func (cat *CategoryControllerImplementation) AddCategory(c *gin.Context) {
	var category *request.Category
	err := c.ShouldBind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, err.Error(), nil))
		return
	}

	isSuccess, msg, err := cat.CategoryService.AddCategory(c.Request.Context(), category)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (cat *CategoryControllerImplementation) UpdateCategory(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	var category *request.Category
	err = c.ShouldBind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, err.Error(), nil))
		return
	}
	category.ID = ID

	isSuccess, msg, err := cat.CategoryService.UpdateCategory(c.Request.Context(), category)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (cat *CategoryControllerImplementation) DeleteCategory(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	isSuccess, msg, err := cat.CategoryService.DeleteCategory(c.Request.Context(), ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (cat *CategoryControllerImplementation) GetCategoryByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	category, msg, err := cat.CategoryService.GetCategoryByID(c.Request.Context(), ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if category != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, category))
		return
	}
}

func (cat *CategoryControllerImplementation) GetCategoryByName(c *gin.Context) {

	name := c.Query("name")

	category, msg, err := cat.CategoryService.GetCategoryByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if category != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, category))
		return
	}
}

func (cat *CategoryControllerImplementation) GetCategories(c *gin.Context) {
	categories, msg, err := cat.CategoryService.GetCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if categories != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, categories))
		return
	}
}
