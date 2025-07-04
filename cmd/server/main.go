package main

import (
	"go-blog/config"
	"go-blog/internal/handler"
	"go-blog/internal/middleware"
	"go-blog/internal/repository"
	"go-blog/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, relying on environment variables.")
	}
	cfg := config.Load()
	db := config.ConnectDB(cfg)

	commentRepo := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepo)
	commentHandler := handler.NewCommentHandler(commentService)

	r := gin.Default()
	r.Use(middleware.Recaptcha())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Go Blog API"})
	})

	api := r.Group("/api/v1")
	commentHandler.RegisterRoutes(api)
	r.Run(":" + cfg.Port)
}
