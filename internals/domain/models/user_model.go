package models

import "time"

// CreateUserRequest - untuk create operation
type CreateUserRequest struct {
    UserName    string `json:"user_name" validate:"required,min=3,max=20"`
    Email       string `json:"email" validate:"email,required"`
    Password    string `json:"password" validate:"required,min=6,max=50"`
}

// UpdateUserRequest - untuk update operation dengan pointer
type UpdateUserRequest struct {
    UserName    string `json:"user_name" validate:"omitempty,min=3,max=20"`
    Email       string `json:"email" validate:"omitempty,email"`
    Password    string `json:"password" validate:"omitempty,min=6,max=50"`
}
type UpdateUserResponse struct {
    ID          string    `json:"id"`
    UserName    string    `json:"user_name"`
    Email       string    `json:"email"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
// UserResponse - untuk response
type UserResponse struct {
    ID          string    `json:"id"`
    UserName    string    `json:"user_name"`
    Email       string    `json:"email"`
    Role        string    `json:"role"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// LoginRequest - untuk authentication
type LoginRequest struct {
    Email       string `json:"email" validate:"email,required"`
    Password    string `json:"password" validate:"required"`
}

// ChangePasswordRequest - untuk change password
type ChangePasswordRequest struct {
    OldPassword     string `json:"old_password" validate:"required"`
    NewPassword     string `json:"new_password" validate:"required,min=6,max=50"`
}


// User - untuk operasi user
type User struct {
    ID        string
    UserName  string   
    Email     string   
    Password  string    
    Role      string
    CreatedAt time.Time 
    UpdatedAt time.Time
}
