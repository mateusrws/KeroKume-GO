package auth

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"kerokume-go/utils"
)

func (s *JwtType) ValidateToken(token string, ctx *gin.Context) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			utils.SendError(ctx, http.StatusUnauthorized, "Invalid signing method")
			return nil, nil
		}
		return []byte(s.SecretKey), nil
	})

	if err != nil {
		utils.SendError(ctx, http.StatusUnauthorized, "Invalid token")
		return false
	}

	return err == nil
}