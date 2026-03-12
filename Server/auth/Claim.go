package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Claim struct{
	Sum uuid.UUID `json: "sum"`
	jwt.StandardClaims
}