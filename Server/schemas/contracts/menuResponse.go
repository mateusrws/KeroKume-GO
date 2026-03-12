package contracts

import "github.com/google/uuid"


type MenuResponse struct {
	Id 	 uuid.UUID `json: "id"`
	Name string 	 `json: "name"`
}