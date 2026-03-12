package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetIdFromJwt(ctx *gin.Context) uuid.UUID {
	authorization := ctx.GetHeader("Authorization")
	token := authorization[len("Bearer "):]
	jwtParts := ConvertJwtInObj(token, ctx)
	restaurantIdStr := jwtParts.Payload.Sum
	restaurantId, err := uuid.Parse(restaurantIdStr)
	if err != nil {
		SendError(ctx, http.StatusUnauthorized, "invalid token")
		return uuid.Nil
	}
	return restaurantId
}
