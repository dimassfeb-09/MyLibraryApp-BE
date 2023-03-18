package controller

import (
	"net/http"

	"github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/request"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/helpers"
	"github.com/dimassfeb-09/MyLibraryApp-BE.git/service"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	AuthRegister(c *gin.Context)
	AuthLogin(c *gin.Context)
}

type AuthControllerImplementation struct {
	AuthService service.AuthService
}

func NewAuthControllerImplementation(authService service.AuthService) AuthController {
	return &AuthControllerImplementation{AuthService: authService}
}

func (a *AuthControllerImplementation) AuthRegister(c *gin.Context) {
	var authRegister request.AuthRegister
	err := c.ShouldBind(&authRegister)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, err.Error(), nil))
		return
	}

	isSuccess, msg, err := a.AuthService.AuthRegister(c.Request.Context(), &authRegister)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusOK, helpers.ToWebResponse("OK", http.StatusOK, "Berhasil Register.", nil))
		return
	}
}

func (a *AuthControllerImplementation) AuthLogin(c *gin.Context) {
	var authLogin request.AuthLogin
	err := c.ShouldBind(&authLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, err.Error(), nil))
		return
	}

	isSuccess, msg, err := a.AuthService.AuthLogin(c.Request.Context(), &authLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ToWebResponse("Bad Request", http.StatusBadRequest, msg, nil))
		return
	}

	if isSuccess {
		c.JSON(http.StatusOK, helpers.ToWebResponse("OK", http.StatusOK, msg, nil))
		return
	}
}
