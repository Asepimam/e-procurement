package usecases

import (
	"context"
	"e-procurement/internals/domain/models"
	"e-procurement/internals/repositories"
	customContext "e-procurement/pkg/context"
	"fmt"
	"log"
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
func (v *VendorUseCase) CreateVendorUsecase(ctx context.Context, vendorReq *models.CreateVendorRequest) (*models.CreateVendorResponse, error) {

	userIdContext,err := customContext.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user ID from context: %w", err)
	}
	if userIdContext == "" {
		return nil, fmt.Errorf("user ID not found in context")
	}

	exitsUser, err := v.userRepository.GetUserByID(ctx, userIdContext)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if exitsUser == nil {
		return nil, fmt.Errorf("user with ID %s does not exist", exitsUser.ID)
	}
	// validate user ready vendor
	userVendorsExists, err := v.vendorRepository.GetVendorByUserID(ctx, exitsUser.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user vendors: %w", err)
	}
	if userVendorsExists != nil {	
		return nil, fmt.Errorf("user already has a vendor")
	}
	vendor, err := v.vendorRepository.CreateVendor(ctx,exitsUser.ID, vendorReq)
	if err != nil {
		return nil, err
	}
	vendorResponse := &models.CreateVendorResponse{
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
func (v *VendorUseCase) GetAllVendors(ctx context.Context, limit, page int) ([]*models.VendorResponse, int, error) {
	
	if limit <= 0 {
		limit = 10 // default limit
	}
	if page <= 0 {
		page = 1 // default offset
	}
	// Calculate offset
	Offset := (page - 1) * limit

	if limit > 100  || Offset > 100 {
		return nil, 0, fmt.Errorf("limit and offset must be less than or equal to 100")
	}
	// Query total count
	count, err := v.vendorRepository.CountVendors(ctx)
	log.Println("Total vendors count:", count)
	if err != nil {
		return nil, 0, err
	}
	log.Println("Limit:", limit, "Offset:", Offset)
	// Get vendors from repository
	vendors, err := v.vendorRepository.GetAllVendors(ctx, limit, Offset )
	log.Println("Vendors retrieved:", vendors)
	if err != nil {
		return nil, 0, err
	}
	var vendorResponses []*models.VendorResponse
	for _, vendor := range vendors {
		vendorResponses = append(vendorResponses, &models.VendorResponse{
			ID:          vendor.ID,
			VendorName:  vendor.VendorName,
			Description: vendor.Description,
			UserID:      vendor.UserID,
			UserName:    vendor.UserName,
			CreatedAt:   vendor.CreatedAt,
			UpdatedAt:   vendor.UpdatedAt,
		})
		
	}
	return vendorResponses, count, nil
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
func (v *VendorUseCase) UpdateVendor(ctx context.Context, id string, vendorReq *models.UpdateVendorRequest) (*models.UpdateVendorResponse, error) {
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

	vendorResponse := &models.UpdateVendorResponse{
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
	