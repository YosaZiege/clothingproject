package models

import "github.com/google/uuid"
import "time"

// Order model represents the orders table in the database
type Order struct {
    ID        uuid.UUID `json:"id" db:"id"`
    UserID    uuid.UUID `json:"user_id" db:"user_id"` // foreign key to Users
    Status    string    `json:"status" db:"status"` // enum: "pending", "shipped", "completed", "canceled"
    TotalPrice float64  `json:"total_price" db:"total_price"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// OrderItem model represents an order item in the order_items table
type OrderItem struct {
    ID        uuid.UUID `json:"id" db:"id"`
    OrderID   uuid.UUID `json:"order_id" db:"order_id"` // foreign key to Orders
    ProductID uuid.UUID `json:"product_id" db:"product_id"` // foreign key to Products
    Quantity  int       `json:"quantity" db:"quantity"`
    Price     float64   `json:"price" db:"price"`
}
