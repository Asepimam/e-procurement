package usecases

import (
	"context"
	"e-procurement/internals/domain/models"
	"e-procurement/internals/repositories"
	customContext "e-procurement/pkg/context"
	"fmt"
)

type ProductUseCase struct {
	productRepository *repositories.ProductRepository
	vendorRepository *repositories.VendorRepository

}
// NewProductUsecase instence 
func NewProductUsecase(productRepo *repositories.ProductRepository, vendor *repositories.VendorRepository ) *ProductUseCase {
	return &ProductUseCase{
		productRepository: productRepo,
		vendorRepository: vendor,
	}
}

// Method for creted new product
func(u *ProductUseCase) CreateProducUsecase(ctx context.Context, productReq *models.CreateProductRequest)(*models.CreateProductResponse, error){
	userID,err := customContext.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user ID from context: %w", err)
	}

	userVendor,err := u.vendorRepository.GetVendorByUserID(ctx, userID)

	if err != nil {
		return nil, fmt.Errorf("failed to get vendor by user ID: %w", err)
	}
	
	if userVendor == nil {
		return nil, fmt.Errorf("vendor with user ID %s does not exist", userID)
	}
	// validate vendor ID and vendor body is user vendor
	if productReq.VendorID != userVendor.ID {
		return nil, fmt.Errorf("vendor ID Canot create product for other vendor")
	}
	product, err := u.productRepository.CreateProduct(ctx, productReq)
	if err != nil {
		return nil, err
	}
	
	reponsesProduct := &models.CreateProductResponse{
		ID:                 	product.ID,
		ProductName: 	  		product.ProductName,
		ProductPrice: 			product.ProductPrice,
		ProductDescription: 	product.ProductDescription,
		ProductCategoryID: 		product.ProductCategoryID,
		VendorID: 				product.VendorID,
		CreatedAt: 				product.CreatedAt,
		UpdatedAt: 				product.UpdatedAt,
	}
	return reponsesProduct, nil
}

// Method to Get All Products
func(u *ProductUseCase) GetAllProducts(ctx context.Context, limit, page int) ([]*models.ResponseProduct, int, error) {
	if limit <= 0 {
		limit = 10 // default limit
	}
	if page <= 0 {
		page = 1 // default offset
	}
	offset := (page - 1) * limit
	if limit > 100  || page > 100 {
		return nil, 0, fmt.Errorf("limit and offset must be between 1 and 100")
	}
	// Query total count
	count, err := u.productRepository.CountProducts(ctx)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to count products: %w", err)
	}

	// Get products from repository
	products, err := u.productRepository.GetAllProducts(ctx, limit, offset)

	if err != nil {
		return nil,0, fmt.Errorf("failed to get all products: %w", err)
	}

	var productsResponse []*models.ResponseProduct
	for _, product := range products {
		productsResponse = append(productsResponse, &models.ResponseProduct{
			ID:                  	product.ID,
			ProductName:        	product.ProductName,
			ProductPrice:       	product.ProductPrice,
			ProductDescription: 	product.ProductDescription,
			ProductCategoryID:  	product.ProductCategoryID,
			ProductCategoryName: 	product.ProductCategoryName,
			VendorID:           	product.VendorID,
			VendorName: 	   		product.VendorName,
			CreatedAt:          	product.CreatedAt,
			UpdatedAt:          	product.UpdatedAt,
		})
	}
	return productsResponse, count, nil
}

// Method to Get Products By Category
func(u *ProductUseCase) GetProductsByCategory(ctx context.Context, category string, limit, offset int) ([]*models.ResponseProduct,int, error) {
	if limit <= 0 {
		limit = 10 // default limit
	}
	if offset < 0 {
		offset = 1 // default offset
	}
	offset = (offset - 1) * limit

	// Get products from repository
	products, err := u.productRepository.GetProductsByCategory(ctx, category, limit, offset)
	if err != nil {
		return nil,0, fmt.Errorf("failed to get products by category: %w", err)
	}

	count, err := u.productRepository.CountProducts(ctx)
	if err != nil {
		return nil,0, fmt.Errorf("failed to count products: %w", err)
	}
	var productsResponse []*models.ResponseProduct
	for _, product := range products {
		productsResponse = append(productsResponse, &models.ResponseProduct{
			ID:                  	product.ID,
			ProductName:        	product.ProductName,
			ProductPrice:       	product.ProductPrice,
			ProductDescription: 	product.ProductDescription,
			ProductCategoryID:  	product.ProductCategoryID,
			ProductCategoryName: 	product.ProductCategoryName,
			VendorID:           	product.VendorID,
			VendorName: 	   		product.VendorName,
			CreatedAt:          	product.CreatedAt,
			UpdatedAt:          	product.UpdatedAt,
		})
	}
	return productsResponse,count,nil
}


// Method to Get Product By ID
func(u *ProductUseCase) GetProductByID(ctx context.Context, id string) (*models.ResponseProduct, error) {
	product, err := u.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get product by ID: %w", err)
	}
	if product == nil {
		return nil, nil // Product not found
	}
	productResponse := &models.ResponseProduct{
		ID:                 	product.ID,
		ProductName:        	product.ProductName,
		ProductPrice:       	product.ProductPrice,
		ProductDescription: 	product.ProductDescription,
		ProductCategoryID:  	product.ProductCategoryID,
		ProductCategoryName: 	product.ProductCategoryName,
		VendorID:           	product.VendorID,
		VendorName: 	   		product.VendorName,
		CreatedAt:          	product.CreatedAt,
		UpdatedAt:          	product.UpdatedAt,
	}
	return productResponse, nil
}

// Method to Update Product
func(u *ProductUseCase) UpdateProduct(ctx context.Context, id string, productReq *models.UpdateProductRequest) (*models.UpdateProductResponse, error) {
	// Check if product exists
	existingProduct, err := u.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get product by ID: %w", err)
	}
	if existingProduct == nil {
		return nil, fmt.Errorf("product with ID %s not found", id)
	}
	
	if productReq.ProductName == "" {
		productReq.ProductName = existingProduct.ProductName
	}
	if productReq.ProductPrice == 0 {
		productReq.ProductPrice = existingProduct.ProductPrice
	}
	if productReq.ProductDescription == "" {
		productReq.ProductDescription = existingProduct.ProductDescription
	}
	if productReq.ProductCategoryID == "" {
		productReq.ProductCategoryID = existingProduct.ProductCategoryID
	}

	// Update product
	updatedProduct, err := u.productRepository.UpdateProduct(ctx, id, productReq)
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}
	productResponse := &models.UpdateProductResponse{
		ID:                  updatedProduct.ID,
		ProductName:         updatedProduct.ProductName,
		ProductPrice:        updatedProduct.ProductPrice,
		ProductDescription:  updatedProduct.ProductDescription,
		ProductCategoryID:   updatedProduct.ProductCategoryID,
		VendorID:            updatedProduct.VendorID,
		CreatedAt:           updatedProduct.CreatedAt,
		UpdatedAt:           updatedProduct.UpdatedAt,
	}

	return productResponse, nil
}

// Method to Delete Product
func(u *ProductUseCase) DeleteProduct(ctx context.Context, id string) error {
	// Check if product exists
	existingProduct, err := u.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get product by ID: %w", err)
	}
	if existingProduct == nil {
		return fmt.Errorf("product with ID %s not found", id)
	}

	// Delete product
	err = u.productRepository.DeleteProduct(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}