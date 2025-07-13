package repositories

import (
	"context"
	"database/sql"
	"e-procurement/internals/domain/models"
	"log"

	sq "github.com/Masterminds/squirrel"
)

type VendorRepository struct {
	db         *sql.DB
	SQLBuilder sq.StatementBuilderType
}

// NewVendorRepository creates a new instance of VendorRepository with the provided database connection.
func NewVendorRepository(db *sql.DB) *VendorRepository {
	return &VendorRepository{
		db:         db,
		SQLBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

// Method to Create New Vendor
// It returns a string representing the ID of the newly created vendor.
// parameters:
// 		ctx: context for request-scoped values and cancellation.
// 		vendorModel : vendor model containing vendor details.
// returns:
// 		VendorResponse: a VendorResponse model containing the details of the created vendor.
// 		errors: if any occurred during the operation.
func (v *VendorRepository) CreateVendor(ctx context.Context, userId string, vendorModel *models.CreateVendorRequest) (*models.Vendor, error) {
	query := v.SQLBuilder.
		Insert("vendors").
		Columns("vendor_name", "description", "user_id").
		Values(vendorModel.VendorName, vendorModel.Description, userId).
		Suffix("RETURNING id, vendor_name, description, user_id, created_at, updated_at")

	row := query.RunWith(v.db).QueryRowContext(ctx)

	var vendorResponse models.Vendor
	err := row.Scan(
		&vendorResponse.ID,
		&vendorResponse.VendorName,
		&vendorResponse.Description,
		&vendorResponse.UserID,
		&vendorResponse.CreatedAt,
		&vendorResponse.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &vendorResponse, nil
}


// Method to Get All Vendors
// It returns a slice of VendorResponse models containing the details of all vendors.
// parameters:
// 		ctx: context for request-scoped values and cancellation.
// 		limit: number of vendors to retrieve.
func (v *VendorRepository) GetAllVendors(ctx context.Context, limit, offset int) ([]*models.Vendor, error) {
	log.Println("Fetching all vendors with limit:", limit, "and offset:", offset, "from database")
	query := v.SQLBuilder.
		Select(
			"v.id", 
			"v.vendor_name", 
			"v.description", 
			"v.user_id", 
			"u.user_name",
			"v.created_at", 
			"v.updated_at",
			).
		From("e_procurement.vendors v").
		LeftJoin("e_procurement.users u ON v.user_id = u.id").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		Suffix("ORDER BY v.created_")

	rows, err := query.RunWith(v.db).QueryContext(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No vendors found
		}
		return nil, err
	}
	defer rows.Close()
	var vendors []*models.Vendor
	for rows.Next() {
		var vendor models.Vendor
		err := rows.Scan(
			&vendor.ID,
			&vendor.VendorName,
			&vendor.Description,
			&vendor.UserID,
			&vendor.UserName,
			&vendor.CreatedAt,
			&vendor.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning vendor row:", err)
			return nil, err
		}
		vendors = append(vendors, &vendor)
	}
	
	return vendors, nil
}
	
// Method to Get Vendor By ID
// It returns a VendorResponse model containing the details of the vendor with the specified ID.
// parameters:
// 		ctx: context for request-scoped values and cancellation.
// 		id: ID of the vendor to retrieve.
// returns:
// 		VendorResponse: a VendorResponse model containing the details of the vendor.
func (v *VendorRepository) GetVendorByID(ctx context.Context, id string) (*models.Vendor, error) {
	query := v.SQLBuilder.
		Select(
			"v.id", 
			"v.vendor_name", 
			"v.description", 
			"v.user_id", 
			"u.user_name",
			"v.created_at", 
			"v.updated_at",
			).
		From("e_procurement.vendors v").
		LeftJoin("e_procurement.users u ON v.user_id = u.id").
		Where(sq.Eq{"v.id": id})

	row := query.RunWith(v.db).QueryRowContext(ctx)

	var vendor models.Vendor
	err := row.Scan(
		&vendor.ID,
		&vendor.VendorName,
		&vendor.Description,
		&vendor.UserID,
		&vendor.UserName,
		&vendor.CreatedAt,
		&vendor.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &vendor, nil
}

// Method to Update Vendor
// It returns a VendorResponse model containing the updated details of the vendor.
// parameters:
// 		ctx: context for request-scoped values and cancellation.
// 		vendorModel: vendor model containing updated vendor details.
// returns:
// 		VendorResponse: a VendorResponse model containing the updated details of the vendor.
// 		errors: if any occurred during the operation.
func (v *VendorRepository) UpdateVendor(ctx context.Context,vendorID string, vendorModel *models.UpdateVendorRequest) (*models.Vendor, error) {
	query := v.SQLBuilder.
		Update("vendors").
		Set("vendor_name", vendorModel.VendorName).
		Set("description", vendorModel.Description).
		Set("user_id", vendorModel.UserID).
		Where(sq.Eq{"id": vendorID}).
		Suffix("RETURNING id, vendor_name, description, user_id, created_at, updated_at")

	row := query.RunWith(v.db).QueryRowContext(ctx)

	var vendorResponse models.Vendor
	err := row.Scan(
		&vendorResponse.ID,
		&vendorResponse.VendorName,
		&vendorResponse.Description,
		&vendorResponse.UserID,
		&vendorResponse.CreatedAt,
		&vendorResponse.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &vendorResponse, nil
}

// Method to Delete Vendor
// It deletes the vendor with the specified ID from the database.
// parameters:
// 		ctx: context for request-scoped values and cancellation.
// 	id: ID of the vendor to delete.
// returns:
// 		errors: if any occurred during the operation.
func (v *VendorRepository) DeleteVendor(ctx context.Context, id string) error {
	query := v.SQLBuilder.
		Delete("vendors").
		Where(sq.Eq{"id": id})
	_, err := query.RunWith(v.db).ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Method to Count Total Vendors
// It returns the total count of vendors in the database.
func (v *VendorRepository) CountVendors(ctx context.Context) (int, error) {
	query := v.SQLBuilder.
		Select("COUNT(*)").
		From("vendors")

	row := query.RunWith(v.db).QueryRowContext(ctx)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Method to Get Vendor By User ID
// It returns a VendorResponse model containing the details of the vendor associated with the specified user ID.
// parameters:
// 		ctx: context for request-scoped values and cancellation.
// 		userID: ID of the user to retrieve the vendor for.
// returns:
// 		VendorResponse: a VendorResponse model containing the details of the vendor.
// 		errors: if any occurred during the operation.
func(v *VendorRepository) GetVendorByUserID(ctx context.Context, userID string) (*models.Vendor, error) {
	query := v.SQLBuilder.
		Select(
			"v.id", 
			"v.vendor_name", 
			"v.description", 
			"v.user_id", 
			"u.user_name",
			"v.created_at", 
			"v.updated_at",
			).
		From("e_procurement.vendors v").
		LeftJoin("e_procurement.users u ON v.user_id = u.id").
		Where(sq.Eq{"v.user_id": userID})

	row := query.RunWith(v.db).QueryRowContext(ctx)

	var vendor models.Vendor
	err := row.Scan(
		&vendor.ID,
		&vendor.VendorName,
		&vendor.Description,
		&vendor.UserID,
		&vendor.UserName,
		&vendor.CreatedAt,
		&vendor.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No vendor found for the user ID
		}
		return nil, err
	}

	return &vendor, nil
}