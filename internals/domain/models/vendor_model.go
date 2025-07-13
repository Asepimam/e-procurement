package models

import (
	"time"
)

// Vendor represents a vendor in the system.
type Vendor struct {
	ID        		string 
	VendorName 		string
	Description 	string
	UserID    		string
	UserName 		string
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}

type CreateVendorRequest struct {
	VendorName 		string `json:"vendor_name" validate:"required"`
	Description 	string `json:"description" validate:"required"`
	UserID     		string `json:"user_id" validate:"required,uuid"`
}

type UpdateVendorRequest struct {
	VendorName 		string `json:"vendor_name" validate:"required"`
	Description 	string `json:"description" validate:"required"`
	UserID     		string `json:"user_id" validate:"required,uuid"`
}

type VendorResponse struct {
	ID          	string    `json:"id"`
	VendorName  	string    `json:"vendor_name"`
	Description 	string    `json:"description"`
	UserID      	string    `json:"user_id"`
	UserName    	string    `json:"user_name"`
	CreatedAt   	time.Time `json:"created_at"`
	UpdatedAt   	time.Time `json:"updated_at"`
}