package models

import "time"

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` 
	FullName     string    `json:"full_name"`
	RoleID       string    `json:"role_id"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
	RoleID   string `json:"role_id" validate:"required"`
}