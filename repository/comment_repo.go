package repositories

import (
	"assignment-4/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	FindAll() ([]models.Comment, error)
	FindByID(id uint) (models.Comment, error)
	Create(comment models.Comment) (models.Comment, error)
	Update(comment models.Comment) (models.Comment, error)
	Delete(id uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) FindAll() ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.Preload("User").Preload("Photo").Find(&comments).Error
	return comments, err
}

func (r *commentRepository) FindByID(id uint) (models.Comment, error) {
	var comment models.Comment
	err := r.db.Preload("User").Preload("Photo").First(&comment, id).Error
	return comment, err
}

func (r *commentRepository) Create(comment models.Comment) (models.Comment, error) {
	err := r.db.Preload("User").Create(&comment).Error
	return comment, err
}

func (r *commentRepository) Update(comment models.Comment) (models.Comment, error) {
	err := r.db.Preload("User").Save(&comment).Error
	return comment, err
}

func (r *commentRepository) Delete(id uint) error {
	err := r.db.Delete(&models.Comment{}, id).Error
	return err
}
