package usecases

import (
	"context"
	"e-procurement/internals/domain/models"
	"e-procurement/internals/repositories"
)

type CategoryUsecase struct {
	CategoryRepository *repositories.CategoryRepository
}

func NewCategoryUsecase(categoryRepository *repositories.CategoryRepository,) *CategoryUsecase{
	return&CategoryUsecase{
		CategoryRepository: categoryRepository,
	}
}

// method for Create New Category
func(u *CategoryUsecase) CreateCategoryUsecase(ctx context.Context, category *models.CreateCategoryRequest) (*models.CreateCategoryResponse, error) {
	categoryResponse, err := u.CategoryRepository.CreateCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return &models.CreateCategoryResponse{
		ID:          categoryResponse.ID,
		Name:        categoryResponse.Name,
		Description: categoryResponse.Description,
		CreatedAt:   categoryResponse.CreatedAt,
		UpdatedAt:   categoryResponse.UpdatedAt,
	}, nil
}