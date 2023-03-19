package controller

import (
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/request"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/helpers"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type WishlistController interface {
	AddWishlist(c *gin.Context)
	UpdateWishlist(c *gin.Context)
	DeleteWishlist(c *gin.Context)
	GetWishlistByID(c *gin.Context)
	GetWishlistByUserID(c *gin.Context)
	GetWishlists(c *gin.Context)
}

type WishlistControllerImplementation struct {
	WishlistService service.WishlistService
}

func NewWishlistControllerImplementation(wishlistService service.WishlistService) WishlistController {
	return &WishlistControllerImplementation{WishlistService: wishlistService}
}

func (w *WishlistControllerImplementation) AddWishlist(c *gin.Context) {
	var wishlist *request.Wishlist
	err := c.ShouldBind(&wishlist)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, err.Error(), nil))
		return
	}

	isSuccess, msg, err := w.WishlistService.AddWishlist(c.Request.Context(), wishlist)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (w *WishlistControllerImplementation) UpdateWishlist(c *gin.Context) {

	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	var wishlist *request.Wishlist
	err = c.ShouldBind(&wishlist)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, err.Error(), nil))
		return
	}
	wishlist.ID = ID

	isSuccess, msg, err := w.WishlistService.UpdateWishlist(c.Request.Context(), wishlist)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (w *WishlistControllerImplementation) DeleteWishlist(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	isSuccess, msg, err := w.WishlistService.DeleteWishlist(c.Request.Context(), ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (w *WishlistControllerImplementation) GetWishlistByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", "asdas"))
		return
	}

	wishlist, msg, err := w.WishlistService.GetWishlistByID(c.Request.Context(), ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, wishlist))
	return
}

func (w *WishlistControllerImplementation) GetWishlistByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", "asdas"))
		return
	}

	wishlist, msg, err := w.WishlistService.GetWishlistByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, wishlist))
	return
}

func (w *WishlistControllerImplementation) GetWishlists(c *gin.Context) {
	wishlists, msg, err := w.WishlistService.GetWishlists(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, wishlists))
	return
}
