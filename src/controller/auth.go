package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunasky-hy/dialy-note-app/src/model"
	"github.com/lunasky-hy/dialy-note-app/src/service"
)

type AuthController struct {
	service service.AuthService
}

func (ac AuthController) Signup(c *gin.Context) {
	var json model.User
	if parse_err := c.ShouldBindJSON(&json); parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": parse_err.Error()})
		return
	}

	token, register_err := ac.service.Register(json)
	
	if register_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": register_err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func (ac AuthController) Signin(c *gin.Context) {
	var json model.User
	if parse_err := c.ShouldBindJSON(&json); parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": parse_err.Error()})
		return
	}

	token, auth_err := ac.service.AuthorizeUser(json)
	
	if auth_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": auth_err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}


func CreateAuthController(service service.AuthService) AuthController {
	controller := AuthController{service: service}
	return controller
}