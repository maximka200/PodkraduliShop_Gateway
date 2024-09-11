package models

type Product struct {
	id          int
	imageUrl    string
	title       string
	description string
	price       int
	currency    rune // validation on db level ($ or ₽)
	discount    int
	productUrl  string
}
