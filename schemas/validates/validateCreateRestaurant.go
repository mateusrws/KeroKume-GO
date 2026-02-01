package validates

import (
	"fmt"

	"kerokume-go/schemas/contracts"
	"kerokume-go/utils"
)

func ValidateCreateRestaurant(r *contracts.RestaurantRequest) error {
	if r.Name == "" && r.Description == "" && r.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Name == "" {
		return utils.ParamIsRequired("name", "string")
	}
	if r.Description == "" {
		return utils.ParamIsRequired("description", "string")
	}
	if r.Password == "" {
		return utils.ParamIsRequired("pass", "string")
	}
	return nil
}
