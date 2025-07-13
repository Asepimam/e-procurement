package usecases

import (
	"context"
	"e-procurement/internals/domain/models"
	"e-procurement/internals/repositories"
	"fmt"
)

type VendorUseCase struct {
	vendorRepository *repositories.VendorRepository
	userRepository	*repositories.UserRepository
}


func NewVendorUseCase(vendorRepo *repositories.VendorRepository,userRepo *repositories.UserRepository) *VendorUseCase {
	return &VendorUseCase{
		vendorRepository: vendorRepo,
		userRepository: userRepo,
	}
}

// method to create a new vendor
func (v *VendorUseCase) CreateVendorUsecase(ctx context.Context, vendorReq *models.CreateVendorRequest) (*models.VendorResponse, error) {
	exitsUser, err := v.userRepository.GetById(ctx, vendorReq.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if exitsUser == nil {
		return nil, fmt.Errorf("user with ID %s does not exist", vendorReq.UserID)
	}
	vendor, err := v.vendorRepository.CreateVendor(ctx, vendorReq)
	if err != nil {
		return nil, err
	}
	vendorResponse := &models.VendorResponse{
	   ID:          vendor.ID,
	   VendorName:  vendor.VendorName,
	   Description: vendor.Description,
	   UserID:      vendor.UserID,
	   CreatedAt:   vendor.CreatedAt,
	   UpdatedAt:   vendor.UpdatedAt,
   }
   return vendorResponse, nil
}

// Method to Get All Vendors
func (v *VendorUseCase) GetAllVendors(ctx context.Context, limit, page int) ([]*models.Vendor, int, error) {
	if limit <= 0 {
		limit = 10 // default limit
	}
	if page <= 0 {
		page = 1 // default offset
	}
	newOffset := (page - 1) * limit

	// Query total count
	count, err := v.vendorRepository.CountVendors(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Get vendors from repository
	vendors, err := v.vendorRepository.GetAllVendors(ctx, limit, newOffset )
	if err != nil {
		return nil, 0, err
	}

	return vendors, count, nil
}

// Method to Get Vendor by ID
func (v *VendorUseCase) GetVendorByID(ctx context.Context, id string) (*models.VendorResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("vendor ID is required")
	}

	vendor, err := v.vendorRepository.GetVendorByID(ctx, id)
	if err != nil {
		return nil, err
	}

	vendorResponse := &models.VendorResponse{
		ID:          vendor.ID,
		VendorName:  vendor.VendorName,
		Description: vendor.Description,
		UserID:      vendor.UserID,
		UserName:    vendor.UserName,
		CreatedAt:   vendor.CreatedAt,
		UpdatedAt:   vendor.UpdatedAt,
	}

	return vendorResponse, nil
}

// Method to Update Vendor
func (v *VendorUseCase) UpdateVendor(ctx context.Context, id string, vendorReq *models.UpdateVendorRequest) (*models.VendorResponse, error) {
	existingVendor, err := v.vendorRepository.GetVendorByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("vendor not found: %w", err)
	}
	if existingVendor == nil {
		return nil, fmt.Errorf("vendor with ID %s does not exist", id)
	}
	if vendorReq.VendorName == "" {
		vendorReq.VendorName = existingVendor.VendorName
	}
	if vendorReq.Description == "" {
		vendorReq.Description = existingVendor.Description
	}
	if vendorReq.UserID == "" {
		vendorReq.UserID = existingVendor.UserID
	}

	updatedVendor, err := v.vendorRepository.UpdateVendor(ctx, id, vendorReq)
	if err != nil {
		return nil, fmt.Errorf("failed to update vendor: %w", err)
	}

	vendorResponse := &models.VendorResponse{
		ID:          updatedVendor.ID,
		VendorName:  updatedVendor.VendorName,
		Description: updatedVendor.Description,
		UserID:      updatedVendor.UserID,
		CreatedAt:   updatedVendor.CreatedAt,
		UpdatedAt:   updatedVendor.UpdatedAt,
	}
	return vendorResponse, nil
}

// method to delete a vendor
func (v *VendorUseCase) DeleteVendor(ctx context.Context, id string) error {
	err := v.vendorRepository.DeleteVendor(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete vendor: %w", err)
	}
	return nil
}
	