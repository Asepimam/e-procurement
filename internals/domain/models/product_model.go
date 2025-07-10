package models

import "time"

type CreateProductRequest struct {
	ProductName        string  `json:"product_name" validate:"required"`
	ProductPrice       float64 `json:"product_price" validate:"required,gt=0"`
	ProductDescription string  `json:"product_description" validate:"required"`
	ProductCategory    string  `json:"product_category" validate:"required"`
}

type UpdateProductRequest struct {
	ProductName        string  `json:"product_name" validate:"omitempty"`
	ProductPrice       float64 `json:"product_price" validate:"omitempty,gt=0"`
	ProductDescription string  `json:"product_description" validate:"omitempty"`
	ProductCategory    string  `json:"product_category" validate:"omitempty"`
}

type ResponseProduct struct {
	ID                 string  `json:"id"`
	ProductName        string  `json:"product_name"`
	ProductPrice       float64 `json:"product_price"`
	ProductDescription string  `json:"product_desc"`
	ProductCategory    string  `json:"product_category"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt		   time.Time `json:"update_at"`	
}

type Product struct {
	ID                 string
	ProductName        string
	ProductPrice       string
	ProductDescription string
	ProductCateory     string
	CreatedAt          time.Time 
	UpdatedAt		   time.Time 	
}