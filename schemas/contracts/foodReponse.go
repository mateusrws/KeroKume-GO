package contracts

import "github.com/google/uuid"

type FoodResponse struct {
	Id 					 uuid.UUID `json: "id"`
	Name 				 string `json: "name"`
	Description  string `json: "description"`
	Price 			 float32 `json: "price"`
	PathImg 		 string `json: "path"`
	FoodCategory string `json: "category"`
	IsAvailable  bool `json: "isAvailable"`
}