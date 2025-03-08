package dto

import "github.com/google/uuid"

type UserResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	PhoneNumber string    `json:"phoneNumber"`
}
