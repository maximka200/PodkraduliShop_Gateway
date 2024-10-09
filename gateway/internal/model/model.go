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

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	AppId    string `json:"app_id`
}

type MainPageReqParam struct {
	Page         int `form:"page"`
	PerPageCount int `form:"pagePerCount"`
}
