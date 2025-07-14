package repositories

import (
	"context"
	"database/sql"
	"e-procurement/internals/domain/models"
	"fmt"

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
		Columns("product_name", "product_price", "product_description", "product_category","vendor_id").
		Values(product.ProductName, product.ProductPrice, product.ProductDescription, product.ProductCategoryID, product.VendorID).
		Suffix("RETURNING id, product_name, product_price, product_description, product_category, vendor_id, created_at, updated_at")
	row := query.RunWith(p.db).QueryRowContext(ctx)

	var productResponse models.Product
	err := row.Scan(
		&productResponse.ID,
		&productResponse.ProductName,
		&productResponse.ProductPrice,
		&productResponse.ProductDescription,
		&productResponse.ProductCategoryID,
		&productResponse.VendorID,
		&productResponse.CreatedAt,
		&productResponse.UpdatedAt,
	)
	if err != nil {
		return nil,err
	}

	return &productResponse, nil
}

// Method to Get All Products
// It returns a slice of ProductResponse models containing the details of all products.
// parameters:
// 		ctx: context for request-scoped values and cancellation.
//      limit: maximum number of products to return.
//      offset: number of products to skip before starting to return results.
// returns:
// 		[]ProductResponse: a slice of ProductResponse models containing the details of all products.
// 		errors: if any occurred during the operation.
func (p *ProductRepository) GetAllProducts(ctx context.Context, limit, offset int) ([]*models.Product, error) {
    query := p.SQLBuilder.
        Select(
            "p.id",
            "p.product_name",
            "p.product_price",
            "p.product_description",
            "p.product_category",
            "c.category_name",
			"p.vendor_id",
			"v.vendor_name",
            "p.created_at",
            "p.updated_at",
        ).
        From("e_procurement.products p").
        LeftJoin("e_procurement.categories c ON p.product_category = c.id").
		LeftJoin("e_procurement.vendors v ON p.vendor_id = v.id").
        Limit(uint64(limit)).
        Offset(uint64(offset))

    rows, err := query.RunWith(p.db).QueryContext(ctx)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []*models.Product
    for rows.Next() {
        var product models.Product
        if err := rows.Scan(
            &product.ID,
            &product.ProductName,
            &product.ProductPrice,
            &product.ProductDescription,
            &product.ProductCategoryID,
            &product.ProductCategoryName,
			&product.VendorID,
			&product.VendorName,
            &product.CreatedAt,
            &product.UpdatedAt,
        ); err != nil {
            return nil, err
        }
        products = append(products, &product)
    }

    return products, nil
}

// Method to Get Product by ID
// It returns a ProductResponse model containing the details of the product with the specified ID.
// parameters:
// 		ctx: context for request-scoped values and cancellation.
//      id: ID of the product to be retrieved.
// returns:
// 		ProductResponse: a ProductResponse model containing the details of the product.	
// 		errors: if any occurred during the operation.
func (p *ProductRepository) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
    query := p.SQLBuilder.
        Select(
            "p.id",
            "p.product_name",
            "p.product_price",
            "p.product_description",
            "p.product_category",
            "c.category_name",
			"p.vendor_id",
			"v.vendor_name",
            "p.created_at",
            "p.updated_at",
        ).
        From("products p").
        LeftJoin("categories c ON p.product_category = c.id").
		LeftJoin("vendors v ON p.vendor_id = v.id").
        Where(sq.Eq{"p.id": id})

    row := query.RunWith(p.db).QueryRowContext(ctx)

    var product models.Product
    err := row.Scan(
        &product.ID,
        &product.ProductName,
        &product.ProductPrice,
        &product.ProductDescription,
        &product.ProductCategoryID,
        &product.ProductCategoryName, 
		&product.VendorID,
		&product.VendorName,
        &product.CreatedAt,
        &product.UpdatedAt,
    )
    if err != nil {
        return nil, fmt.Errorf("failed to get product by ID: %w", err)
    }

    return &product, nil
}

// Method to Update Product by ID
// It returns a ProductResponse model containing the updated details of the product.
// parameters:
// 		ctx: context for request-scoped values and cancellation.
//      id: ID of the product to be updated.
//      product: pointer to UpdateProductRequest model containing the updated product details.
// returns:
// 		ProductResponse: a ProductResponse model containing the updated details of the product.
// 		errors: if any occurred during the operation.
func(p *ProductRepository) UpdateProduct(ctx context.Context, id string, product *models.UpdateProductRequest) (*models.Product, error) {
	query := p.SQLBuilder.
		Update("products").
		Set("product_name", product.ProductName).
		Set("product_price", product.ProductPrice).
		Set("product_description", product.ProductDescription).
		Set("product_category", product.ProductCategoryID).
		Where(sq.Eq{"id": id}).
		Suffix("RETURNING id, product_name, product_price, product_description, product_category, vendor_id, created_at, updated_at")

	row := query.RunWith(p.db).QueryRowContext(ctx)

	var updatedProduct models.Product
	err := row.Scan(
		&updatedProduct.ID,
		&updatedProduct.ProductName,
		&updatedProduct.ProductPrice,
		&updatedProduct.ProductDescription,
		&updatedProduct.ProductCategoryID,
		&updatedProduct.VendorID,
		&updatedProduct.CreatedAt,
		&updatedProduct.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &updatedProduct, nil
}

// Method to Delete Product by ID
// It returns an error if any occurred during the operation.
func(p *ProductRepository) DeleteProduct(ctx context.Context, id string) error {
	query := p.SQLBuilder.
		Delete("products").
		Where(sq.Eq{"id": id})
	_, err := query.RunWith(p.db).ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Method to Get Products by Category
// It returns a slice of ProductResponse models containing the details of products in the specified category.
// parameters:
// 		ctx: context for request-scoped values and cancellation.
//      category: category of the products to be retrieved.
// returns:
// 		[]Product: a slice of Product as category models containing the details of products in
// 		the specified category.
// 		errors: if any occurred during the operation.
func(p *ProductRepository) GetProductsByCategory(ctx context.Context, category string, limit, offset int) ([]*models.Product, error) {
	query := p.SQLBuilder.
		 Select(
            "p.id",
            "p.product_name",
            "p.product_price",
            "p.product_description",
            "p.product_category",
            "c.category_name",
            "p.created_at",
            "p.updated_at",
        ).
        From("e_procurement.products p").
        Join("e_procurement.categories c ON p.product_category = c.id").
        OrderBy("p.created_at DESC").
		Where(sq.Eq{"c.category_name": category}).
		Limit(uint64(limit)).
		Offset(uint64(offset))

	rows, err := query.RunWith(p.db).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(
			&product.ID,
			&product.ProductName,
			&product.ProductPrice,
			&product.ProductDescription,
			&product.ProductCategoryID,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}


	return products, nil
}

// Method Count Products
// It returns the total number of products in the database.
// parameters:
// 		ctx: context for request-scoped values and cancellation.
// returns:
// 		int: total number of products.
func(p *ProductRepository) CountProducts(ctx context.Context) (int, error) {
	query := p.SQLBuilder.
		Select("COUNT(*)").
		From("products")

	row := query.RunWith(p.db).QueryRowContext(ctx)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}