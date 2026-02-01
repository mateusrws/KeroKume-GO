package contracts

type RestaurantRequest struct {
	Name string `json: "name"`
	Description string `json: "description"`
	Password string `json: "pass"`
}