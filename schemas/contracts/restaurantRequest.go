package contracts

type RestaurantRequest struct {
	Name 				string `json: "name"`
	Description string `json: "description"`
	Email 			string `json: "email"`
	Password 		string `json: "pass"`
}