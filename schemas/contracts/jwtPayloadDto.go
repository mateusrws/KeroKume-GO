package contracts

type JwtPayloadDto struct {
	Sum string `json:"Sum"`
	Exp int64  `json:"exp"`
	Iss string `json:"iss"`
	Iat int64  `json:"iat"`
}
