package model

type GetProductResponse struct {
	Id          string
	ImageURL    string
	Title       string
	Description string
	Price       int64
	Currency    string // unmarshal in Unicode
	Discount    uint32
	ProductURL  string
}
