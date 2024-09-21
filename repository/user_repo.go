package repositories

import (
	"assignment-4/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	FindByUsernameOrEmail(username string, email string) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) FindByUsernameOrEmail(username string, email string) (models.User, error) {
	user := models.User{}
	err := r.db.Where("username = ? OR email = ?", username, email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
