package contracts


type ConvertJwtDto struct {
	Header    JwtHeaderDto 	`json:"header"`
	Payload   JwtPayloadDto `json:"payload"`
	Signature string 				`json:"signature"`
}
