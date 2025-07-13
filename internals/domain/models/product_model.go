package models

import "time"

type CreateProductRequest struct {
	ProductName        		string  `json:"product_name" validate:"required"`
	ProductPrice       		float64 `json:"product_price" validate:"required,gt=0"`
	ProductDescription 		string  `json:"product_description" validate:"required"`
	ProductCategoryID    	string  `json:"product_category_id" validate:"required"`
}

type UpdateProductRequest struct {
	ProductName        		string  `json:"product_name" validate:"omitempty"`
	ProductPrice       		float64 `json:"product_price" validate:"omitempty,gt=0"`
	ProductDescription 		string  `json:"product_description" validate:"omitempty"`
	ProductCategoryID    	string  `json:"product_category_id" validate:"omitempty"`
}

type ResponseProduct struct {
	ID                 		string  `json:"id"`
	ProductName        		string  `json:"product_name"`
	ProductPrice       		float64 `json:"product_price"`
	ProductDescription 		string  `json:"product_desc"`
	ProductCategoryID    	string  `json:"product_category_id"`
	ProductCategoryName  	string  `json:"product_category_name"`
	VendorID           		string  `json:"vendor_id"`
	VendorName         		string  `json:"vendor_name"`
	CreatedAt          		time.Time `json:"created_at"`
	UpdatedAt		   		time.Time `json:"update_at"`	
}

type Product struct {
	ID                 		string
	ProductName        		string
	ProductPrice       		float64
	ProductDescription 		string
	ProductCategoryID  		string
	ProductCategoryName 	string
	VendorID           		string
	VendorName         		string
	CreatedAt          		time.Time 
	UpdatedAt		   		time.Time 	
}