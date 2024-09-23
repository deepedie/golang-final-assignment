package services

import (
	"assignment-4/models"
	repositories "assignment-4/repository"
)

type PhotoService interface {
	GetAllPhotos() ([]models.Photo, error)
	GetPhotoByID(id uint) (models.Photo, error)
	CreatePhoto(photo models.Photo) (models.Photo, error)
	UpdatePhoto(photo models.Photo) (models.Photo, error)
	DeletePhoto(id uint) error
}

type photoService struct {
	repo repositories.PhotoRepository
}

func NewPhotoService(repo repositories.PhotoRepository) PhotoService {
	return &photoService{repo: repo}
}

func (s *photoService) GetAllPhotos() ([]models.Photo, error) {
	return s.repo.FindAll()
}

func (s *photoService) GetPhotoByID(id uint) (models.Photo, error) {
	return s.repo.FindByID(id)
}

func (s *photoService) CreatePhoto(photo models.Photo) (models.Photo, error) {
	photo, err := s.repo.Create(photo)
	if err != nil {
		return models.Photo{}, err
	}
	return s.repo.FindByID(photo.ID)
}

func (s *photoService) UpdatePhoto(photo models.Photo) (models.Photo, error) {
	photo, err := s.repo.Update(photo)
	if err != nil {
		return models.Photo{}, err
	}
	return s.repo.FindByID(photo.ID)
}

func (s *photoService) DeletePhoto(id uint) error {
	return s.repo.Delete(id)
}
