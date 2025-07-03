package service

import (
	"go-blog/internal/model"
	"go-blog/internal/repository"
)

type CommentService interface {
	AddComment(comment *model.Comment) error
	GetComments(postID uint) ([]model.Comment, error)
}

type commentService struct {
	repo repository.CommentRepository
}

func NewCommentService(r repository.CommentRepository) CommentService {
	return &commentService{r}
}

func (s *commentService) AddComment(c *model.Comment) error {
	return s.repo.Create(c)
}

func (s *commentService) GetComments(postID uint) ([]model.Comment, error) {
	return s.repo.ListByPost(postID)
}
