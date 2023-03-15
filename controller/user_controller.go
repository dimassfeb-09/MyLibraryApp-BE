package controller

import (
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/request"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserController interface {
	AddUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	GetUserByNPM(c *gin.Context)
	GetUserByEmail(c *gin.Context)
}

type UserControllerImplementation struct {
	UserService service.UserService
}

func NewUserControllerImplementation(userService service.UserService) UserController {
	return &UserControllerImplementation{UserService: userService}
}

func (u *UserControllerImplementation) AddUser(c *gin.Context) {
	var user request.AddUser
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	isSuccess, msg, err := u.UserService.AddUser(c.Request.Context(), &user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, map[string]any{
			"status": "Bad Request",
			"code":   http.StatusBadRequest,
			"msg":    msg,
			"data":   nil,
		})
		return
	}

	if isSuccess {
		c.JSON(http.StatusOK, map[string]any{
			"status": "OK",
			"code":   http.StatusOK,
			"msg":    msg,
			"data":   nil,
		})
		return
	}
}

func (u *UserControllerImplementation) UpdateUser(c *gin.Context) {
	// TODO
}

func (u *UserControllerImplementation) DeleteUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControllerImplementation) GetUserByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	userByID, msg, err := u.UserService.GetUserByID(c.Request.Context(), ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"status": "Bad Request",
			"code":   http.StatusBadRequest,
			"msg":    msg,
			"data":   nil,
		})
		return
	}
	if userByID != nil {
		c.JSON(http.StatusOK, map[string]any{
			"status": "OK",
			"code":   http.StatusOK,
			"msg":    msg,
			"data":   userByID,
		})
		return
	}
}

func (u *UserControllerImplementation) GetUserByNPM(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControllerImplementation) GetUserByEmail(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
