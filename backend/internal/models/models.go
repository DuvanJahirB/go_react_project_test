package models

type User struct {
    Email    string `json:"email" bson:"email" validate:"required,email"`
    Password string `json:"password" bson:"password" validate:"required,min=6"`
    Name     string `json:"name" bson:"name" validate:"required"`
}

type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

// MessageResponse represents a generic success message response.
type MessageResponse struct {
    Message string `json:"message" example:"Operation successful"`
}

// ErrorResponse represents a generic error message response.
type ErrorResponse struct {
    Error string `json:"error" example:"Invalid input"`
}

// TokenResponse represents the response containing a JWT token.
type TokenResponse struct {
    Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// ProfileResponse represents the user profile information.
type ProfileResponse struct {
    Email string `json:"email" example:"user@example.com"`
    Name  string `json:"name" example:"User Name"`
}