package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	ginmiddleware "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimiter() gin.HandlerFunc {
	rate, _ := limiter.NewRateFromFormatted("5-M") // 5 requests per minute
	store := memory.NewStore()
	instance := limiter.New(store, rate)
	return ginmiddleware.NewMiddleware(instance)
}
