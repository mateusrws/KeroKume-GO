package schemas

import (
	"github.com/google/uuid"
)

type Food struct{
	BaseModel
	Name 				 string
	Description  string
	Price 			 float32
	PathImg 		 string
	FoodCategory string
	IsAvailable  bool
	MenuId 			 uuid.UUID
}