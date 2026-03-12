package validates

import (
	"fmt"

	"github.com/google/uuid"

	"kerokume-go/schemas/contracts"
	"kerokume-go/utils"
)

func ValidateCreateMenu(m *contracts.MenuRequest) error {
	if m.Name == "" && m.RestaurantId == uuid.Nil{
		return fmt.Errorf("request body is empty or malformed")
	}
	if m.Name == "" {
		return utils.ParamIsRequired("name", "string")
	}
	if m.RestaurantId == uuid.Nil {
		return utils.ParamIsRequired("restaurantId", "UUID")
	}
	return nil
}
