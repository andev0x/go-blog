package repository

import (
	"go-blog/internal/model"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment *model.Comment) error
	ListByPost(postID uint) ([]model.Comment, error)
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepo{db}
}

func (r *commentRepo) Create(comment *model.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepo) ListByPost(postID uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := r.db.Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}
