package repositories

import (
	"context"
	"database/sql"
	"e-procurement/internals/domain/models"

	sq "github.com/Masterminds/squirrel"
)

type ProductRepository struct {
	db 		*sql.DB
	SQLBuilder 	sq.StatementBuilderType
}

// NewProductUseCase creates a new instance of ProductUseCase with the provided database connection.
func NewProductUseCase(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db:         db,
		SQLBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}


// Method to Create New Product Vendor
// It returns a string representing the ID of the newly created product vendor.
//
// parameters:
// 		ctx: context for request-scoped values and cancellation.
// 		productModel : product model containing product details.
// returns:
// 		ProductResponse: a ProductResponse model containing the details of the created product.
// 		errors: if any occurred during the operation.
func(p *ProductRepository) CreateProduct(ctx context.Context, product *models.CreateProductRequest ) (*models.Product, error) {
	query := p.SQLBuilder.
		Insert("products").
		Columns("product_name", "product_price", "product_description", "product_category").
		Values(product.ProductName, product.ProductPrice, product.ProductDescription, product.ProductCategory).
		Suffix("RETURNING id, product_name, product_price, product_description, product_category")
	row := query.RunWith(p.db).QueryRowContext(ctx)

	var productResponse models.Product
	err := row.Scan(
		&productResponse.ID,
		&productResponse.ProductName,
		&productResponse.ProductPrice,
		&productResponse.ProductDescription,
		&productResponse.ProductCateory,
		&productResponse.CreatedAt,
		&productResponse.UpdatedAt,
	)
	if err != nil {
		return nil,err
	}

	return &productResponse, nil
}