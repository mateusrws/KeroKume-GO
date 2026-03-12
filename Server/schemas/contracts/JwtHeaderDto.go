package contracts

type JwtHeaderDto struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
