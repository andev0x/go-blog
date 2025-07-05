package main

import (
	"go-blog/config"
	"go-blog/internal/handler"
	"go-blog/internal/middleware"
	"go-blog/internal/repository"
	"go-blog/internal/service"
	"log"

	"github.com/gin-contrib/cors"
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

	// CORS configuration - more permissive for testing
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://andev0x.github.io", "http://localhost:3000", "*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.Use(middleware.Recaptcha())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Go Blog API"})
	})

	// Test endpoint to verify routing
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Test endpoint working"})
	})

	api := r.Group("/api/v1")
	commentHandler.RegisterRoutes(api)

	log.Printf("Server starting on port %s", cfg.Port)
	r.Run(":" + cfg.Port)
}
