package contracts

import "github.com/google/uuid"


type FoodRequest struct {
	Name 					string 					`json:"name"`
	Description 	string 					`json:"description"`
	Price 				float32 				`json:"price"`
	FoodCategory 	string 					`json:"category"`
	MenuId 				uuid.UUID 			`json:"menu"`
}