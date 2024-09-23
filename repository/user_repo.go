package repositories

import (
	"assignment-4/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	FindByUsernameOrEmail(username string, email string) (models.User, error)
	ExistsByEmail(email string) bool
	ExistsByUsername(username string) bool
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

func (r *userRepository) ExistsByEmail(email string) bool {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	return result.RowsAffected > 0
}

func (r *userRepository) ExistsByUsername(username string) bool {
	var user models.User
	result := r.db.Where("username = ?", username).First(&user)
	return result.RowsAffected > 0
}
