package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"inception/api/internal/global"
	"inception/api/internal/utils/jwt"
)

func GetClaims(c *gin.Context) (*jwt.BaseClaims, error) {
	anyClaims, exists := c.Get(Claims)
	global.Log.Info("anyClaims", anyClaims)
	if !exists {
		return nil, errors.New("not exists user claims")
	}
	claimsObj, suc := anyClaims.(*jwt.BaseClaims)
	if !suc {
		return nil, errors.New("claims type can't translate to jwt.BaseClaims")
	}
	return claimsObj, nil
}
