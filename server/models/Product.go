package models

import (
	"time"
)


type Product struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name" validate:"required"`
	Description string  `json:"description" db:"description" validate:"required"`
	Price       float64 `json:"price" db:"price" validate:"required,min=0"` // Ensure price is not negative
	Stock       string  `json:"stock" db:"stock"`
	Sold 		int 	`json:"sold" db:"sold" validate:"min=0"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
// type Color struct {
// 	Hex string `json:"hex"`
// 	Name string `json:"name"`
// 	ImageUrl string `json:"image_url"`
// }

// type Review struct {
// 	Rating string `json:"rating"`
// 	Comment string `json:"comment"`
// 	Username string `json:"username"`
// }
// type ProductPageData struct {
// 	Product Product `json:"product"`
// 	Images []string `json:"images"`
// 	Colors []Color	`json:"colors"`
// 	Sizes []string  `json:"sizes"`
// 	Reviews []Review `json:"review"`
// 	IsInCart bool    `json:"isincart"`
// 	IsFavorite bool  `json:"isfavorite"`
// }
