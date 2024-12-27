package models

import "github.com/google/uuid"
import "time"

// User model represents the users table in the database
type User struct {
    ID        uuid.UUID `json:"id" db:"id"`
    Name      string    `json:"name" db:"name"`
    Email     string    `json:"email" db:"email"`
    Password  string    `json:"password" db:"password_hash"`
	ImageUrl  string 	`json:"image_url" db:"image_url"`
	Role      string    `json:"role" db:"role"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
