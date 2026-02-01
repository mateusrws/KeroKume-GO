package contracts

type RestaurantResponse struct {
	id uint `json: "id"`
	name string `json: "name"`
	description string `json: "description"`
}