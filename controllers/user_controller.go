package controllers

import (
	"assignment-4/helpers"
	"assignment-4/models"
	"assignment-4/services"
	"assignment-4/utils"
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
		switch e := err.(type) {
		case *helpers.UniqueViolationError:
			c.JSON(http.StatusConflict, gin.H{
				"statusCode": e.StatusCode,
				"error":      "Conflict",
				"message":    e.Error(),
			})
			return
		case *helpers.ValidationError:
			c.JSON(http.StatusBadRequest, gin.H{
				"statusCode": e.StatusCode,
				"error":      "Validation Error",
				"message":    e.Error(),
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"statusCode": http.StatusInternalServerError,
				"error":      "Internal Server Error",
				"message":    "An error occurred during registration.",
			})
			return
		}
	}

	// newUser, err := ctrl.UserService.Register(user)
	// if err != nil {
	// 	if strings.Contains(err.Error(), "required") || strings.Contains(err.Error(), "minstringlength") {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error":   "Validation Error",
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	if err.Error() == "email must be unique" {
	// 		c.JSON(http.StatusConflict, gin.H{
	// 			"error":   "Conflict",
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	if err.Error() == "username must be unique" {
	// 		c.JSON(http.StatusConflict, gin.H{
	// 			"error":   "Conflict",
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error":   "Internal Server Error",
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

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
