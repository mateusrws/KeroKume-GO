package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kerokume-go/services"
	"kerokume-go/utils"
)

func IsMySelf() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the token and return a string
		const bearerSchema = "Bearer "
		
		header := ctx.GetHeader("Authorization")
		token := header[len(bearerSchema):]

		jwtParts := utils.ConvertJwtInObj(token, ctx)

		idSended := jwtParts.Payload.Sum

		if services.IfExistInDb(idSended, ctx) == false {
			utils.SendError(ctx, http.StatusUnauthorized, "Is not permited execute this action in other restaurants")
		}

		

	}
}
