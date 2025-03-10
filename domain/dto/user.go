package dto

import "github.com/google/uuid"

type UserResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	PhoneNumber string    `json:"phoneNumber"`
}

type UpdateUserRequest struct {
	Name            string  `json:"name:" validate:"required"`
	Username        string  `json:"username:" validate:"required"`
	Email           string  `json:"email" validate:"required"`
	PhoneNumber     string  `json:"phoneNumber" validate:"required"`
	Password        *string `json:"password,omitempty"`
	ConfirmPassword *string `json:"confirmPassword,omitempty"`
	RoleID          uint
}
