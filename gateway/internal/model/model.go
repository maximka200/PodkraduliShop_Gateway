package model

type GetProductResponse struct {
	Id          string `json:"id"`
	ImageURL    string `json:"imageURL"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Currency    string `json:"currency"` // unmarshal in Unicode
	Discount    uint32 `json:"discount"`
	ProductURL  string `json:"producURL"`
}
