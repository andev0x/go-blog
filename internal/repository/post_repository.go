package repository

import (
	"go-blog/internal/model"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db}
}

func (r *PostRepository) GetPostIDBySlug(slug string) (uint, error) {
	var post model.Post
	if err := r.db.Where("slug = ?", slug).First(&post).Error; err != nil {
		return 0, err
	}
	return post.ID, nil
}
