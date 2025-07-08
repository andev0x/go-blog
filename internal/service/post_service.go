package service

import "go-blog/internal/repository"

// PostService provides post-related business logic
// (expand as needed for your app)
type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo}
}

func (s *PostService) GetPostIDBySlug(slug string) (uint, error) {
	return s.repo.GetPostIDBySlug(slug)
}
