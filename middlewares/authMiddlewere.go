package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kerokume-go/services"
	"kerokume-go/utils"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const bearerSchema = "Bearer "
		header := ctx.GetHeader("Authorization")
		if header == "" {
			utils.SendError(ctx, http.StatusUnauthorized, "Not found token")
		}
		token := header[len(bearerSchema):]

		if !services.JwtServices().ValidateToken(token, ctx) {
			utils.SendError(ctx, http.StatusUnauthorized, "Unathorized")
		}
	}
}
