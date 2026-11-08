package model

type Good struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID *int    `json:"category_id"`
}

type CreateGoodInput struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID *int    `json:"category_id"`
}
