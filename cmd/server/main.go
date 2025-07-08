package main

import (
	"go-blog/config"
	"go-blog/internal/handler"
	"go-blog/internal/middleware"
	"go-blog/internal/model"
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

	// Auto-migrate database tables
	if err := db.AutoMigrate(&model.Comment{}, &model.Post{}); err != nil {
		log.Printf("Warning: Auto-migration failed: %v", err)
		log.Println("Attempting to create table manually...")

		// Try to create the table manually with proper PostgreSQL syntax
		sql := `
		CREATE TABLE IF NOT EXISTS comments (
			id SERIAL PRIMARY KEY,
			post_id INTEGER NOT NULL,
			name TEXT NOT NULL,
			content TEXT NOT NULL,
			rating INTEGER NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`

		if err := db.Exec(sql).Error; err != nil {
			log.Printf("Manual table creation failed: %v", err)
			log.Fatalf("Database setup failed. Please check your database configuration.")
		} else {
			log.Println("Table created manually successfully")
		}
	} else {
		log.Println("Database migration completed successfully")
	}

	commentRepo := repository.NewCommentRepository(db)
	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	commentService := service.NewCommentService(commentRepo)
	commentHandler := handler.NewCommentHandler(commentService, postService)

	r := gin.Default()

	// CORS configuration - more permissive for testing
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://andev0x.github.io", "http://localhost:3000", "http://localhost:5173", "http://localhost:5174", "http://localhost:5175", "*"},
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
