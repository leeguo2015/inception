package utils

import (
	"github.com/gin-gonic/gin"
	"inception/api/internal/utils/jwt"
)

func GinContextClaims(c *gin.Context) (*jwt.BaseClaims, bool) {
	claimsAny, exists := c.Get(jwt.UserClaims)
	if !exists {
		return nil, false
	}
	b, ok := claimsAny.(*jwt.BaseClaims)
	if !ok {
		return nil, false
	}
	return b, true
}
