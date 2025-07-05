package handler

import (
	"go-blog/internal/model"
	"go-blog/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	service service.CommentService
}

func NewCommentHandler(s service.CommentService) *CommentHandler {
	return &CommentHandler{s}
}

func (h *CommentHandler) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/comments")
	r.GET("/", h.ListAll)
	r.POST("/", h.Create)
	r.GET("/post/:id", h.ListByPost)

	// Add post-specific routes that frontend expects
	posts := rg.Group("/posts")
	posts.GET("/:slug/comments", h.ListBySlug)
	posts.POST("/:slug/comments", h.CreateBySlug)
	posts.GET("/:slug/ratings", h.GetRatingsBySlug)
	posts.POST("/:slug/ratings", h.CreateRatingBySlug)
}

func (h *CommentHandler) ListAll(c *gin.Context) {
	// Simple endpoint to list all comments
	comments := []model.Comment{} // For now, return empty array
	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) Create(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format: " + err.Error()})
		return
	}

	// Validate required fields
	if comment.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}
	if comment.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content is required"})
		return
	}

	if err := h.service.AddComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, comment)
}

func (h *CommentHandler) ListByPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}
	comments, err := h.service.GetComments(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

// New endpoint for frontend: /posts/:slug/comments
func (h *CommentHandler) ListBySlug(c *gin.Context) {
	_ = c.Param("slug") // TODO: implement slug-based logic
	// For now, return empty array. You can implement slug-based logic later
	comments := []model.Comment{}
	c.JSON(http.StatusOK, comments)
}

// New endpoint for frontend: /posts/:slug/comments (POST)
func (h *CommentHandler) CreateBySlug(c *gin.Context) {
	_ = c.Param("slug") // TODO: implement slug-based logic

	var commentData struct {
		Author  string `json:"author"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&commentData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format: " + err.Error()})
		return
	}

	// Validate required fields
	if commentData.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Author is required"})
		return
	}
	if commentData.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content is required"})
		return
	}

	// Create comment (for now, using slug as post_id)
	comment := model.Comment{
		Name:    commentData.Author,
		Content: commentData.Content,
		PostID:  1, // Placeholder - you'll need to map slug to post_id
	}

	if err := h.service.AddComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return in frontend expected format
	response := gin.H{
		"id":        comment.ID,
		"author":    comment.Name,
		"content":   comment.Content,
		"timestamp": comment.CreatedAt,
	}
	c.JSON(http.StatusCreated, response)
}

// New endpoint for frontend: /posts/:slug/ratings
func (h *CommentHandler) GetRatingsBySlug(c *gin.Context) {
	_ = c.Param("slug") // TODO: implement slug-based logic
	// For now, return empty ratings. You can implement rating aggregation later
	ratings := gin.H{
		"average": 0,
		"total":   0,
		"ratings": []int{},
	}
	c.JSON(http.StatusOK, ratings)
}

// New endpoint for frontend: /posts/:slug/ratings (POST)
func (h *CommentHandler) CreateRatingBySlug(c *gin.Context) {
	_ = c.Param("slug") // TODO: implement slug-based logic

	var ratingData struct {
		Value int `json:"value"`
	}

	if err := c.ShouldBindJSON(&ratingData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format: " + err.Error()})
		return
	}

	// Validate rating value
	if ratingData.Value < 1 || ratingData.Value > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rating must be between 1 and 5"})
		return
	}

	// Create rating (for now, using slug as post_id)
	comment := model.Comment{
		Name:    "Anonymous",
		Content: "Rating",
		PostID:  1, // Placeholder - you'll need to map slug to post_id
		Rating:  ratingData.Value,
	}

	if err := h.service.AddComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return in frontend expected format
	response := gin.H{
		"id":        comment.ID,
		"value":     comment.Rating,
		"timestamp": comment.CreatedAt,
	}
	c.JSON(http.StatusCreated, response)
}
