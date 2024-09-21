package controllers

import (
	"assignment-4/helpers"
	"assignment-4/models"
	"assignment-4/services"
	"assignment-4/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

// NewUserController creates a new instance of UserController
func NewUserController(service services.UserService) *UserController {
	return &UserController{
		UserService: service,
	}
}

func (ctrl *UserController) Register(c *gin.Context) {
	contentType := helpers.GetContentType(c)

	user := models.User{}
	var err error

	if contentType == utils.AppJSON {
		err = c.ShouldBindJSON(&user)
	} else {
		err = c.ShouldBind(&user)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	newUser, err := ctrl.UserService.Register(user)
	if err != nil {
		fmt.Println("ERROR", err)
		if err.Error() == "email must be unique" {
			c.JSON(http.StatusConflict, gin.H{
				"error":   "Conflict",
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "An error occurred during registration.",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       newUser.ID,
		"username": newUser.Username,
		"email":    newUser.Email,
		"age":      newUser.Age,
	})
}

func (ctrl *UserController) Login(c *gin.Context) {
	contentType := helpers.GetContentType(c)

	user := models.User{}
	var err error

	if contentType == utils.AppJSON {
		err = c.ShouldBindJSON(&user)
	} else {
		err = c.ShouldBind(&user)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	token, userExist, err := ctrl.UserService.Login(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":       userExist.ID,
			"username": userExist.Username,
			"email":    userExist.Email,
			"age":      userExist.Age,
		},
	})
}
