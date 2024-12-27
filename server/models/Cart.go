package models

import "github.com/google/uuid"
import "time"

// Cart model represents the cart table in the database
type Cart struct {
    ID        uuid.UUID `json:"id" db:"id"`
    UserID    uuid.UUID `json:"user_id" db:"user_id"` // foreign key to Users
    CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// CartItem model represents a cart item in the cart_items table
type CartItem struct {
    ID        uuid.UUID `json:"id" db:"id"`
    CartID    uuid.UUID `json:"cart_id" db:"cart_id"` // foreign key to Cart
    ProductID uuid.UUID `json:"product_id" db:"product_id"` // foreign key to Products
    Quantity  int       `json:"quantity" db:"quantity"`
}
