package usecases

import (
	"context"
	"e-procurement/internals/domain/models"
	"e-procurement/internals/repositories"
)

type ProductUseCase struct {
	productRepository *repositories.ProductRepository

}
// NewProductUsecase instence 
func NewProductUsecase(productRepo *repositories.ProductRepository,) *ProductUseCase {
	return &ProductUseCase{
		productRepository: productRepo,
	}
}

// Method for creted new product
func(u *ProductUseCase) CreateProducUsecase(ctx context.Context, productReq *models.CreateProductRequest)(*models.ResponseProduct, error){
	product, err := u.productRepository.CreateProduct(ctx, productReq)
	if err != nil {
		return nil, err
	}
	return &models.ResponseProduct{
		ID:                 product.ID,
		ProductName: 	  product.ProductName,
		ProductPrice: 		productReq.ProductPrice,
		ProductDescription: product.ProductDescription,
		ProductCategory: product.ProductCateory,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	},nil
}