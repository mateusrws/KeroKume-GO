package services

import (
	"kerokume-go/auth"
)

func JwtServices() *auth.JwtType {
	return &auth.JwtType{
		SecretKey: "secret-key",
		Issuer:    "kerokume-api",
	}
}
