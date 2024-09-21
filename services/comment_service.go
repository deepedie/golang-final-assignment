package services

import (
	"assignment-4/models"
	repositories "assignment-4/repository"
)

type CommentService interface {
	GetAllComments() ([]models.Comment, error)
	GetCommentByID(id uint) (models.Comment, error)
	CreateComment(comment models.Comment) (models.Comment, error)
	UpdateComment(comment models.Comment) (models.Comment, error)
	DeleteComment(id uint) error
}

type commentService struct {
	repo repositories.CommentRepository
}

func NewCommentService(repo repositories.CommentRepository) CommentService {
	return &commentService{repo: repo}
}

func (s *commentService) GetAllComments() ([]models.Comment, error) {
	return s.repo.FindAll()
}

func (s *commentService) GetCommentByID(id uint) (models.Comment, error) {
	return s.repo.FindByID(id)
}

func (s *commentService) CreateComment(comment models.Comment) (models.Comment, error) {
	return s.repo.Create(comment)
}

func (s *commentService) UpdateComment(comment models.Comment) (models.Comment, error) {
	return s.repo.Update(comment)
}

func (s *commentService) DeleteComment(id uint) error {
	return s.repo.Delete(id)
}
