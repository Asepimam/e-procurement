package repositories

import (
	"context"
	"database/sql"
	"e-procurement/internals/domain/models"

	sq "github.com/Masterminds/squirrel"
)

type CategoryRepository struct {
	db *sql.DB
	SQLBuilder sq.StatementBuilderType
}

// NewCategoryRepository creates a new instance of CategoryRepository with the provided database connection.
func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db:         db,
		SQLBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

// Method to Create New Category
// paramters:
// 		ctx: context for the database operation
// 		category: pointer to CreateCategoryRequest model containing the category details
// returns:
// 		*models.Category: pointer to the created category model
// 		error: error if any occurred during the operation
func(c *CategoryRepository) CreateCategory(ctx context.Context, category *models.CreateCategoryRequest) (*models.Category, error) {
	query := c.SQLBuilder.
		Insert("categories").
		Columns("category_name", "descriptions").
		Values(category.Name,category.Description).
		Suffix("RETURNING id, category_name, descriptions, created_at, updated_at")

	row := query.RunWith(c.db).QueryRowContext(ctx)
	var categoryResponse models.Category
	err := row.Scan(
		&categoryResponse.ID,
		&categoryResponse.Name,
		&categoryResponse.Description,
		&categoryResponse.CreatedAt,
		&categoryResponse.UpdatedAt,)
	if err != nil {
		return nil,err
	}
	return &categoryResponse,nil
}

// Method to Get All Categories
// parameters:
// 		ctx: context for the database operation
// returns:
// 		[]*models.Category: slice of pointers to Category models
// 		error: error if any occurred during the operation
func(c *CategoryRepository) GetAllCategories(ctx context.Context,limit,offset int) ([]*models.Category, error) {
	query := c.SQLBuilder.
		Select("id", "category_name", "descriptions", "created_at", "updated_at").
		From("categories").
		Limit(uint64(limit)).
		Offset(uint64(offset))


	rows, err := query.RunWith(c.db).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(
			&category.ID, 
			&category.Name,
			&category.Description,
			&category.CreatedAt,
			&category.UpdatedAt,
		); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

// Method to Get Category by ID
// parameters:
// 		ctx: context for the database operation
// 		id: ID of the category to be retrieved
// returns:
// 		*models.Category: pointer to the Category model if found
// 		error: error if any occurred during the operation
func(c *CategoryRepository) GetCategoryByID(ctx context.Context, id string) (*models.Category, error) {
	query := c.SQLBuilder.
		Select("id", "category_name", "descriptions", "created_at", "updated_at").
		From("categories").
		Where(sq.Eq{"id": id}).
		Limit(1)

	row := query.RunWith(c.db).QueryRowContext(ctx)
	var category models.Category
	err := row.Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No category found
		}
		return nil, err
	}
	return &category, nil
}

// Method to Update Category
// parameters:
// 		ctx: context for the database operation
// 		category: pointer to UpdateCategoryRequest model containing the updated category details
// returns:
// 		*models.Category: pointer to the updated Category model
// 		error: error if any occurred during the operation
func(c *CategoryRepository) UpdateCategory(ctx context.Context,id string, category *models.UpdateCategoryRequest) (*models.Category, error) {
	query := c.SQLBuilder.
		Update("categories").
		Set("category_name", category.Name).
		Set("descriptions", category.Description).
		Where(sq.Eq{"id": id}).
		Suffix("RETURNING id, category_name, descriptions, created_at, updated_at")

	row := query.RunWith(c.db).QueryRowContext(ctx)
	var updatedCategory models.Category
	err := row.Scan(
		&updatedCategory.ID,
		&updatedCategory.Name,
		&updatedCategory.Description,
		&updatedCategory.CreatedAt,
		&updatedCategory.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No category found to update
		}
		return nil, err
	}
	return &updatedCategory, nil
}

// Method to Delete Category
// parameters:
// 		ctx: context for the database operation
// 		id: ID of the category to be deleted
// returns:
// 		error: error if any occurred during the operation
func(c *CategoryRepository) DeleteCategory(ctx context.Context, id string) error {
	query := c.SQLBuilder.
		Delete("categories").
		Where(sq.Eq{"id": id})

	result, err := query.RunWith(c.db).ExecContext(ctx)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows // No category found to delete
	}
	return nil
}

func(c *CategoryRepository) CountAllCategories(ctx context.Context) (int, error) {
	query := c.SQLBuilder.
		Select("COUNT(*)").
		From("categories")

	row := query.RunWith(c.db).QueryRowContext(ctx)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}