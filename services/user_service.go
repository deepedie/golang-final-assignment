package services

import (
	"assignment-4/helpers"
	"assignment-4/models"
	repositories "assignment-4/repository"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type UserService interface {
	Register(user models.User) (models.User, error)
	Login(user models.User) (string, models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(user models.User) (models.User, error) {
	createdUser, err := s.userRepo.Create(user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return models.User{}, errors.New("email must be unique")
		}
		return models.User{}, err // Return other errors as-is
	}
	return createdUser, nil
}

func (s *userService) Login(user models.User) (string, models.User, error) {
	existingUser, err := s.userRepo.FindByUsernameOrEmail(user.Username, user.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", models.User{}, errors.New("user not found")
		}
		return "", models.User{}, err
	}

	if !helpers.ComparePassword([]byte(existingUser.Password), []byte(user.Password)) {
		return "", models.User{}, errors.New("wrong password")
	}

	token := helpers.GenerateToken(existingUser.ID, existingUser.Email)
	return token, existingUser, nil
}
