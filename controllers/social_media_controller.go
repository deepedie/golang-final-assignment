package controllers

import (
	"assignment-4/models"
	"assignment-4/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	socialMediaService services.SocialMediaService
}

func NewSocialMediaController(socialMediaService services.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{socialMediaService: socialMediaService}
}

func (sc *SocialMediaController) GetAllSocialMedias(c *gin.Context) {
	socialMedias, err := sc.socialMediaService.GetAllSocialMedias()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, socialMedias)
}

func (sc *SocialMediaController) GetSocialMediaByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media ID"})
		return
	}

	socialMedia, err := sc.socialMediaService.GetSocialMediaByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

func (sc *SocialMediaController) CreateSocialMedia(c *gin.Context) {
	// Get the user ID from JWT token
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var socialMedia models.SocialMedia

	// Bind JSON to socialMedia struct
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the UserID for the social media
	socialMedia.UserID = userID

	// Create the social media record
	createdSocialMedia, err := sc.socialMediaService.CreateSocialMedia(socialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the response with the User relationship preloaded
	c.JSON(http.StatusCreated, createdSocialMedia)
}

func (sc *SocialMediaController) UpdateSocialMedia(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var socialMedia models.SocialMedia

	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	socialMedia.ID = uint(id)

	existingSM, err := sc.socialMediaService.GetSocialMediaByID(socialMedia.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	if existingSM.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this photo"})
		return
	}

	socialMedia.UserID = userID

	updatedSocialMedia, err := sc.socialMediaService.UpdateSocialMedia(socialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedSocialMedia)
}

func (sc *SocialMediaController) DeleteSocialMedia(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media ID"})
		return
	}

	err = sc.socialMediaService.DeleteSocialMedia(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Social media deleted successfully"})
}
