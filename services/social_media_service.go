package services

import (
	"assignment-4/models"
	repositories "assignment-4/repository"
)

type SocialMediaService interface {
	GetAllSocialMedias() ([]models.SocialMedia, error)
	GetSocialMediaByID(id uint) (models.SocialMedia, error)
	CreateSocialMedia(socialMedia models.SocialMedia) (models.SocialMedia, error)
	UpdateSocialMedia(socialMedia models.SocialMedia) (models.SocialMedia, error)
	DeleteSocialMedia(id uint) error
}

type socialMediaService struct {
	repo repositories.SocialMediaRepository
}

func NewSocialMediaService(repo repositories.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{repo: repo}
}

func (s *socialMediaService) GetAllSocialMedias() ([]models.SocialMedia, error) {
	return s.repo.FindAll()
}

func (s *socialMediaService) GetSocialMediaByID(id uint) (models.SocialMedia, error) {
	return s.repo.FindByID(id)
}

func (s *socialMediaService) CreateSocialMedia(socialMedia models.SocialMedia) (models.SocialMedia, error) {
	return s.repo.Create(socialMedia)
}

func (s *socialMediaService) UpdateSocialMedia(socialMedia models.SocialMedia) (models.SocialMedia, error) {
	return s.repo.Update(socialMedia)
}

func (s *socialMediaService) DeleteSocialMedia(id uint) error {
	return s.repo.Delete(id)
}
