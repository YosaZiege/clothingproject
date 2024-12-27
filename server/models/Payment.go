package models

import "github.com/google/uuid"
import "time"

// Payment model represents the payments table in the database
type Payment struct {
    ID          uuid.UUID `json:"id" db:"id"`
    OrderID     uuid.UUID `json:"order_id" db:"order_id"` // foreign key to Orders
    PaymentMethod string   `json:"payment_method" db:"payment_method"` // payment method, e.g., "credit_card", "paypal"
    Status      string    `json:"status" db:"status"` // enum: "success", "failed", "pending"
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
