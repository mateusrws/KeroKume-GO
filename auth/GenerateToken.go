package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"kerokume-go/utils"
)

func (s *JwtType) GenerateToken(id uuid.UUID, ctx *gin.Context) (string, error) {
	claim := Claim{
		Sum: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		utils.SendError(ctx, 500, "Generate Token error")
		return "", err
	}
	return t, nil
}
