package controllers

import (
	"assignment-4/models"
	"assignment-4/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	photoService services.PhotoService
}

func NewPhotoController(service services.PhotoService) *PhotoController {
	return &PhotoController{
		photoService: service,
	}
}

func (pc *PhotoController) GetAllPhotos(c *gin.Context) {
	photos, err := pc.photoService.GetAllPhotos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, photos)
}

func (pc *PhotoController) GetPhotoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	photo, err := pc.photoService.GetPhotoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}
	c.JSON(http.StatusOK, photo)
}

func (pc *PhotoController) CreatePhoto(c *gin.Context) {
	var photo models.Photo

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	photo.UserID = userID

	createdPhoto, err := pc.photoService.CreatePhoto(photo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdPhoto)
}

func (pc *PhotoController) UpdatePhoto(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var photo models.Photo

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	photo.ID = uint(id)

	existingPhoto, err := pc.photoService.GetPhotoByID(photo.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	if existingPhoto.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this photo"})
		return
	}

	photo.UserID = userID

	updatedPhoto, err := pc.photoService.UpdatePhoto(photo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedPhoto)
}

func (pc *PhotoController) DeletePhoto(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := pc.photoService.DeletePhoto(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
