package repositories

import (
	"assignment-4/models"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	FindAll() ([]models.SocialMedia, error)
	FindByID(id uint) (models.SocialMedia, error)
	Create(socialMedia models.SocialMedia) (models.SocialMedia, error)
	Update(socialMedia models.SocialMedia) (models.SocialMedia, error)
	Delete(id uint) error
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepository{db: db}
}

func (r *socialMediaRepository) FindAll() ([]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia
	err := r.db.Preload("User").Find(&socialMedias).Error
	return socialMedias, err
}

func (r *socialMediaRepository) FindByID(id uint) (models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	err := r.db.Preload("User").First(&socialMedia, id).Error
	return socialMedia, err
}

func (r *socialMediaRepository) Create(socialMedia models.SocialMedia) (models.SocialMedia, error) {
	err := r.db.Preload("User").Create(&socialMedia).Error
	return socialMedia, err
}

func (r *socialMediaRepository) Update(socialMedia models.SocialMedia) (models.SocialMedia, error) {
	err := r.db.Save(&socialMedia).Error
	return socialMedia, err
}

func (r *socialMediaRepository) Delete(id uint) error {
	err := r.db.Delete(&models.SocialMedia{}, id).Error
	return err
}
