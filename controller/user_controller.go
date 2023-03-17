package controller

import (
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/request"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/helpers"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController interface {
	AddUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	GetUserByNPM(c *gin.Context)
}

type UserControllerImplementation struct {
	UserService service.UserService
}

func NewUserControllerImplementation(userService service.UserService) UserController {
	return &UserControllerImplementation{UserService: userService}
}

func (u *UserControllerImplementation) AddUser(c *gin.Context) {
	var user request.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, err.Error(), nil))
		return
	}

	isSuccess, msg, err := u.UserService.AddUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (u *UserControllerImplementation) UpdateUser(c *gin.Context) {
	var user request.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, err.Error(), nil))
		return
	}

	isSuccess, msg, err := u.UserService.UpdateUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}

func (u *UserControllerImplementation) DeleteUser(c *gin.Context) {
	var user struct {
		ID int `json:"id"`
	}

	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	isSuccess, msg, err := u.UserService.DeleteUser(c.Request.Context(), user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusOK, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}

}

func (u *UserControllerImplementation) GetUserByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	userByID, msg, err := u.UserService.GetUserByID(c.Request.Context(), ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}
	if userByID != nil {
		c.JSON(http.StatusOK, helpers.ToWebResponse("OK", http.StatusOK, msg, userByID))
		return
	}
}

func (u *UserControllerImplementation) GetUserByNPM(c *gin.Context) {
	NPM := c.Query("npm")
	userByNPM, msg, err := u.UserService.GetUserByNPM(c.Request.Context(), NPM)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}
	if userByNPM != nil {
		c.JSON(http.StatusOK, helpers.ToWebResponse("OK", http.StatusOK, msg, userByNPM))
		return
	}
}
