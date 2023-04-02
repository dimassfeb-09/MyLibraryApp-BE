package controller

import (
	"net/http"
	"strconv"

	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/request"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/helpers"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GenreController interface {
	AddGenre(c *gin.Context)
	UpdateGenre(c *gin.Context)
	DeleteGenre(c *gin.Context)
	GetGenreByID(c *gin.Context)
	GetGenreByCategoryID(c *gin.Context)
	GetGenres(c *gin.Context)
}

type GenreControllerImplementation struct {
	GenreService service.GenreService
}

func NewGenreControllerImplementation(genreService service.GenreService) GenreController {
	return &GenreControllerImplementation{GenreService: genreService}
}

func (g *GenreControllerImplementation) AddGenre(c *gin.Context) {
	var genre *request.Genre
	err := c.ShouldBind(&genre)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, err.Error(), nil))
		return
	}

	isSuccess, msg, err := g.GenreService.AddGenre(c.Request.Context(), genre)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusOK, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (g *GenreControllerImplementation) UpdateGenre(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	var genre *request.Genre
	err = c.ShouldBind(&genre)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, err.Error(), nil))
		return
	}
	genre.ID = ID

	isSuccess, msg, err := g.GenreService.UpdateGenre(c.Request.Context(), genre)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helpers.ToWebResponse("Not Found", http.StatusNotFound, msg, nil))
			return
		}
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusOK, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (g *GenreControllerImplementation) DeleteGenre(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	isSuccess, msg, err := g.GenreService.DeleteGenre(c.Request.Context(), ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helpers.ToWebResponse("Not Found", http.StatusNotFound, msg, nil))
			return
		}
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusOK, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (g *GenreControllerImplementation) GetGenreByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	result, msg, err := g.GenreService.GetGenreByID(c.Request.Context(), ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helpers.ToWebResponse("Not Found", http.StatusNotFound, msg, nil))
			return
		}
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	c.JSON(http.StatusOK, helpers.ToWebResponse("OK", http.StatusOK, msg, result))
	return
}

func (g *GenreControllerImplementation) GetGenreByCategoryID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("category_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	result, msg, err := g.GenreService.GetGenreByCategoryID(c.Request.Context(), ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helpers.ToWebResponse("Not Found", http.StatusNotFound, msg, nil))
			return
		}
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	c.JSON(http.StatusOK, helpers.ToWebResponse("OK", http.StatusOK, msg, result))
	return
}

func (g *GenreControllerImplementation) GetGenres(c *gin.Context) {

	results, msg, err := g.GenreService.GetGenres(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	c.JSON(http.StatusOK, helpers.ToWebResponse("OK", http.StatusOK, msg, results))
	return
}
