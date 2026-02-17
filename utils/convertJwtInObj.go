package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"kerokume-go/schemas/contracts"
)

func ConvertJwtInObj(token string, ctx *gin.Context) contracts.ConvertJwtDto {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		SendError(ctx, http.StatusInternalServerError, "Auth middleware contract broken")
		return contracts.ConvertJwtDto{}
	}

	header := parts[0]
	payload := parts[1]
	signature := parts[2]

	payloadJSON, err := jwt.DecodeSegment(payload)
	if err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return contracts.ConvertJwtDto{}
	}

	headerJSON, err := jwt.DecodeSegment(header)
	if err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return contracts.ConvertJwtDto{}
	}

	signatureJSON, err := jwt.DecodeSegment(signature)
	if err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return contracts.ConvertJwtDto{}
	}

	var payloadObj contracts.JwtPayloadDto
	err = json.Unmarshal(payloadJSON, &payloadObj)
	if err != nil {
		SendError(ctx, 500, err.Error())
		return contracts.ConvertJwtDto{}
	}

	var headerObj contracts.JwtHeaderDto
	err = json.Unmarshal(headerJSON, &headerObj)
	if err != nil {
		SendError(ctx, 500, err.Error())
		return contracts.ConvertJwtDto{}
	}

	result := contracts.ConvertJwtDto{
		Header:    headerObj,
		Payload:   payloadObj,
		Signature: string(signatureJSON),
	}

	return result
}
	