package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/voideus/golang-mux-rest/service"
)

type FirebaseAuthMiddleware struct {
	service.FirebaseService
}

func NewAuthMiddleware(service service.FirebaseService) FirebaseAuthMiddleware {
	return FirebaseAuthMiddleware{
		service,
	}
}

func (s FirebaseAuthMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		idToken := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))

		token, err := s.VerifyToken(idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set("ID", token.UID)
		c.Set("role", token.Claims["role"])
		c.Set("has_access", token.Claims["has_access"])
		c.Set("profile_update", token.Claims["profile_updated"])
		c.Next()
	}
}
