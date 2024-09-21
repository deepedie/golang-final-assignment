package repositories

import (
	"assignment-4/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	FindAll() ([]models.Photo, error)
	FindByID(id uint) (models.Photo, error)
	Create(photo models.Photo) (models.Photo, error)
	Update(photo models.Photo) (models.Photo, error)
	Delete(id uint) error
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{db: db}
}

func (r *photoRepository) FindAll() ([]models.Photo, error) {
	var photos []models.Photo
	err := r.db.Preload("User").Find(&photos).Error
	return photos, err
}

func (r *photoRepository) FindByID(id uint) (models.Photo, error) {
	var photo models.Photo
	err := r.db.Preload("User").First(&photo, id).Error
	return photo, err
}

func (r *photoRepository) Create(photo models.Photo) (models.Photo, error) {
	err := r.db.Preload("User").Create(&photo).Error
	return photo, err
}

func (r *photoRepository) Update(photo models.Photo) (models.Photo, error) {
	err := r.db.Preload("User").Save(&photo).Error
	return photo, err
}

func (r *photoRepository) Delete(id uint) error {
	err := r.db.Delete(&models.Photo{}, id).Error
	return err
}
