package middleware

import (
	"net/http"
	"strings"

	error "github.com/astertechs-dev/bizportal-go-backend/domain/error"
	"github.com/astertechs-dev/bizportal-go-backend/internal/token_util"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := token_util.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := token_util.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, error.ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, error.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, error.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
