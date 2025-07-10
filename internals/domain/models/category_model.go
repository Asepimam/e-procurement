package models

import "time"

type CreateCategoryRequest struct {
	Name        string 		`json:"name" validate:"required"`
	Description string 		`json:"description" validate:"required"`
}

type CreateCategoryResponse struct {
	ID          string  		`json:"id"`
	Name        string 		`json:"name"`
	Description string 		`json:"description"`
	CreatedAt   time.Time 	`json:"created_at"`
	UpdatedAt   time.Time 	`json:"updated_at"`
}

type UpdateCategoryRequest struct {
	ID          string  		`json:"id" validate:"required"`
	Name        string 		`json:"name" validate:"required"`
	Description string 		`json:"description" validate:"required"`
}


type Category struct {
	ID			string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt	time.Time
}