package controller

import (
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/request"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/helpers"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RatingController interface {
	AddRating(c *gin.Context)
	UpdateRating(c *gin.Context)
	DeleteRating(c *gin.Context)
	GetRatingByBookID(c *gin.Context)
	GetRatingByID(c *gin.Context)
}

type RatingControllerImplementation struct {
	RatingService service.RatingService
}

func NewRatingControllerImplementation(ratingService service.RatingService) RatingController {
	return &RatingControllerImplementation{
		RatingService: ratingService,
	}
}

func (r *RatingControllerImplementation) AddRating(c *gin.Context) {
	var rating *request.Rating
	err := c.ShouldBind(&rating)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, err.Error(), nil))
		return
	}

	isSuccess, msg, err := r.RatingService.AddRating(c.Request.Context(), rating)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helpers.ToWebResponse("NOT FOUND", http.StatusNotFound, msg, nil))
			return
		}
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (r *RatingControllerImplementation) UpdateRating(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	var rating *request.Rating
	err = c.ShouldBind(&rating)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, err.Error(), nil))
		return
	}
	rating.ID = ID

	isSuccess, msg, err := r.RatingService.UpdateRating(c.Request.Context(), rating)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helpers.ToWebResponse("NOT FOUND", http.StatusNotFound, msg, nil))
			return
		}
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (r *RatingControllerImplementation) DeleteRating(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	isSuccess, msg, err := r.RatingService.DeleteRating(c.Request.Context(), ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helpers.ToWebResponse("NOT FOUND", http.StatusNotFound, msg, nil))
			return
		}
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (r *RatingControllerImplementation) GetRatingByBookID(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Query("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	results, msg, err := r.RatingService.GetRatingByBookID(c.Request.Context(), bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, results))
	return
}

func (r *RatingControllerImplementation) GetRatingByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	result, msg, err := r.RatingService.GetRatingByID(c.Request.Context(), ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helpers.ToWebResponse("NOT FOUND", http.StatusNotFound, msg, nil))
			return
		}
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if result != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, result))
		return
	}
}
