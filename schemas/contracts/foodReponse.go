package contracts

type FoodResponse struct {
	Name string `json: "name"`
	Description string `json: "description"`
	Price float32 `json: "price"`
	PathImg string `json: "path"`
	FoodCategory string `json: "category"`
	IsAvailable bool `json: "isAvailable"`
}