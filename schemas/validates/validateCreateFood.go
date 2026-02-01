package validates

import (
	"fmt"

	"github.com/google/uuid"

	"kerokume-go/schemas/contracts"
	"kerokume-go/utils"
)

func ValidateFoodRequest(f *contracts.FoodRequest) error {
	if f.Name == "" && f.Description == "" && f.Price <= 0 && f.FoodCategory == "" && f.MenuId == uuid.Nil {
		return fmt.Errorf("request body is empty or malformed")
	}
	if f.Name == "" {
		return utils.ParamIsRequired("name", "string")
	}
	if f.Description == "" {
		return utils.ParamIsRequired("description", "string")
	}
	if f.Price <= 0 {
		return utils.ParamIsRequired("price", "float")
	}
	if f.FoodCategory == "" {
		return utils.ParamIsRequired("category", "string")
	}
	if f.MenuId == uuid.Nil {
		return utils.ParamIsRequired("menu", "uuid")
	}
	return nil
}
