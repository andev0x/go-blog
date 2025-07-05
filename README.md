# Go Blog Backend

A Go-based backend API for the andev0x tech blog, built with Gin, GORM, and SQLite/PostgreSQL.

## Features

- RESTful API for comments and ratings
- SQLite for local development, PostgreSQL for production
- CORS enabled for frontend integration
- Rate limiting middleware
- Clean architecture with handlers, services, and repositories

## Quick Start

### Prerequisites

- Go 1.24.3 or higher
- SQLite (for local development)
- PostgreSQL (for production)

### Local Development

1. **Clone and navigate to the backend directory:**
   ```bash
   cd go-blog
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables:**
   ```bash
   # Create .env file (already exists)
   PORT=8080
   DB_PATH=blog.db
   ```

4. **Initialize database:**
   ```bash
   # The database will be created automatically when you first run the server
   # Or you can manually create it using the migration:
   sqlite3 blog.db < migrations/create_comments_table.sql
   ```

5. **Run the server:**
   ```bash
   go run cmd/server/main.go
   ```

The server will start on `http://localhost:8080`

### Production Deployment

For production, set the `DATABASE_URL` environment variable:

```bash
DATABASE_URL="postgres://username:password@localhost:5432/blogdb"
PORT=8080
```

## API Endpoints

### Comments

- `GET /api/v1/posts/:slug/comments` - Get comments for a post
- `POST /api/v1/posts/:slug/comments` - Add a comment to a post

**Comment POST body:**
```json
{
  "author": "Your Name",
  "content": "Your comment text"
}
```

### Ratings

- `GET /api/v1/posts/:slug/ratings` - Get ratings for a post
- `POST /api/v1/posts/:slug/ratings` - Rate a post

**Rating POST body:**
```json
{
  "value": 5
}
```

### Health Check

- `GET /test` - Test endpoint to verify the server is running

## Project Structure

```
go-blog/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── config/
│   └── config.go            # Configuration management
├── internal/
│   ├── handler/
│   │   └── comment_handler.go # HTTP request handlers
│   ├── middleware/
│   │   ├── recaptcha.go     # ReCAPTCHA middleware
│   │   └── rate_limiter.go  # Rate limiting middleware
│   ├── model/
│   │   └── comment.go       # Data models
│   ├── repository/
│   │   └── comment_repository.go # Database operations
│   └── service/
│       └── comment_service.go # Business logic
├── migrations/
│   └── create_comments_table.sql # Database schema
├── go.mod                   # Go module dependencies
├── go.sum                   # Dependency checksums
└── .env                     # Environment variables
```

## Database Schema

The `comments` table stores both comments and ratings:

```sql
CREATE TABLE comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    content TEXT NOT NULL,
    rating INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

## Frontend Integration

The frontend is configured to connect to this backend via the `VITE_API_BASE_URL` environment variable. When the backend is unavailable, the frontend falls back to mock data.

## Development Notes

- The backend currently uses placeholder logic for slug-to-post-id mapping
- Comments and ratings are stored in the same table for simplicity
- CORS is configured to allow requests from `localhost:3000` and `andev0x.github.io`
- Rate limiting is set to 5 requests per minute per IP

## Next Steps

1. Implement proper slug-to-post-id mapping
2. Add authentication for comment moderation
3. Implement proper rating aggregation
4. Add database migrations for schema changes
5. Add comprehensive logging and monitoring
