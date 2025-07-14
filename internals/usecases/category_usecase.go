package usecases

import (
	"context"
	"e-procurement/internals/domain/models"
	"e-procurement/internals/repositories"
	"fmt"
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
func(u *CategoryUsecase) CreateCategoryUsecase(ctx context.Context, category *models.CreateCategoryRequest) (*models.CategoryResponse, error) {
	categoryResponse, err := u.CategoryRepository.CreateCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return &models.CategoryResponse{
		ID:          categoryResponse.ID,
		Name:        categoryResponse.Name,
		Description: categoryResponse.Description,
		CreatedAt:   categoryResponse.CreatedAt,
		UpdatedAt:   categoryResponse.UpdatedAt,
	}, nil
}

// method for Get All Categories
func(u *CategoryUsecase) GetAllCategoriesUsecase(ctx context.Context, limit, page int) ([]*models.CategoryResponse,int, error) {
	if limit <= 0 {
		limit = 10 // default limit
	}
	if page <= 0 {
		page = 1 // default offset
	}
	offset := (page - 1) * limit
	// Query total count
	count, err := u.CategoryRepository.CountAllCategories(ctx)
	if err != nil {
		return nil, 0, err
	}
	categories, err := u.CategoryRepository.GetAllCategories(ctx, limit, offset)
	if err != nil {
		return nil,0, err
	}
	var categoryResponses []*models.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, &models.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
		})
	}
	return categoryResponses, count,nil
}

// method for Get Category By ID
func(u *CategoryUsecase) GetCategoryByIDUsecase(ctx context.Context, id string) (*models.CategoryResponse, error) {
	category, err := u.CategoryRepository.GetCategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}
	categoryResponse := &models.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
	return categoryResponse, nil
}

// method for Update Category
func(u *CategoryUsecase) UpdateCategoryUsecase(ctx context.Context, id string, category *models.UpdateCategoryRequest) (*models.CategoryResponse, error) {
	exists,err := u.CategoryRepository.GetCategoryByID(ctx, id)
	// check if the category exists with error repository
	if err != nil {
		return nil, err
	}
	// check if the category exists
	if category.Name == "" {
		category.Name = exists.Name
	}
	if category.Description == "" {
		category.Description = exists.Description
	}
	// if the category does not exist, return nil
	if exists == nil {
		return nil,fmt.Errorf("category with id %s not found", id)
	}
	updatedCategory, err := u.CategoryRepository.UpdateCategory(ctx,id, category)
	if err != nil {
		return nil, err
	}
	categoryResponse := &models.CategoryResponse{
		ID:          updatedCategory.ID,
		Name:        updatedCategory.Name,
		Description: updatedCategory.Description,
		CreatedAt:   updatedCategory.CreatedAt,
		UpdatedAt:   updatedCategory.UpdatedAt,
	}
	return categoryResponse, nil
}

// method for Delete Category
func(u *CategoryUsecase) DeleteCategoryUsecase(ctx context.Context, id string) error {
	err := u.CategoryRepository.DeleteCategory(ctx, id)
	if err != nil {
		return err
	}
	return nil
}