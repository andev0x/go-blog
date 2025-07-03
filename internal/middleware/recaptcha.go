package middleware

import (
	"github.com/gin-gonic/gin"
)

func Recaptcha() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Dummy middleware logic for ReCAPTCHA
		c.Next()
	}
}
