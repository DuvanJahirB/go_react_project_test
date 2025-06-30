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