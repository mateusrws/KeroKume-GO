package contracts

import "github.com/google/uuid"

type MenuRequest struct{
	Name 					string 					`json: "name"`
	RestaurantId 	uuid.UUID 			`json: "description"`
}