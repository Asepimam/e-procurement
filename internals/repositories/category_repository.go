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
// 
func(c *CategoryRepository) CreateCategory(ctx context.Context, category *models.CreateCategoryRequest) (*models.Category, error) {
	query := c.SQLBuilder.
		Insert("categories").
		Columns("name", "description").
		Values(category.Name,category.Description).
		Suffix("RETURNING id, name, description, created_at, updated_at")

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