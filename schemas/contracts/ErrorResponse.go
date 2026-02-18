package contracts

type ErrorResponse struct {
	Message 	string `json: "message`
	ErrorCode string `json: errorCode`
}