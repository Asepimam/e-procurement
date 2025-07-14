package repositories

import (
	"context"
	"database/sql"
	"e-procurement/internals/domain/models"

	sq "github.com/Masterminds/squirrel"
)

type UserRepository struct {
	db         *sql.DB
	SQKBuilder sq.StatementBuilderType
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db:         db,
		SQKBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}


	// Method to fetch all users from the database.
	// It returns a slice of UserResponse modelss containing user details.
	// 
	// parameters:
	// 		ctx: context for request-scoped values and cancellation.
	// 		limit: maximum number of users to fetch.
	// 		offset: number of users to skip before starting to fetch.	
	// 
	// returns: 
	// 		UserResponse: a slice of UserResponse containing user details.
	//  	error: if any occurred during the operation.
func (r *UserRepository) GetAll(ctx context.Context, limit, offset int) ([]models.UserResponse, error) {
	query := r.SQKBuilder.
		Select("id, user_name, email, role, created_at, updated_at").
		From("users").
		Limit(uint64(limit)).
		Offset(uint64(offset))

	// Execute query
	rows, err := query.RunWith(r.db).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserResponse
	for rows.Next() {
		var user models.UserResponse
		if err := rows.Scan(&user.ID, &user.UserName, &user.Email,&user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	
	return users, nil
}


	// Method to fetch a user by ID from the database.
	// It returns a UserResponse models containing user details.
	// 
	// parameters:
	// 		ctx: context for request-scoped values and cancellation.
	// 		id: the unique identifier of the user to be fetched.
	// 
	// returns: 
	// 		UserResponse : a pointer to UserResponse containing user details.
	// 		error : if any occurred during the operation.
func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	query := r.SQKBuilder.
		Select("id, user_name, email, password, role").
		From("e_procurement.users").
		Where(sq.Eq{"id": id}).
		Limit(1)

	row := query.RunWith(r.db).QueryRowContext(ctx)
	var user models.User

	err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.Role)
	if err != nil {
		
		if err == sql.ErrNoRows {
			return nil, nil // No user found, return nil
		}
		return nil, err // Return error if any other issue occurred
	}
	
	return &user, nil
}


	// Method to create a new user in the database.
	// It inserts the user details into the "users" table and returns the created user response.
	// 
	// parameter:
	// 		ctx: context for request-scoped values and cancellation.
	// 		user: the user models containing the details to be inserted.
	// 
	// returns:
	// 		userResponse: a pointer to UserResponse containing the created user details.
	// 		error: an error if any occurred during the operation.
func (r *UserRepository) Create(ctx context.Context, user *models.CreateUserRequest) (*models.User, error) {
	query := r.SQKBuilder.
		Insert("users").
		Columns("user_name", "email", "password").
		Values(user.UserName, user.Email, user.Password).
		Suffix("RETURNING id, user_name, email, role, created_at, updated_at")

	row := query.RunWith(r.db).QueryRowContext(ctx)

	var userResponse models.User
	err := row.Scan(&userResponse.ID, &userResponse.UserName, &userResponse.Email,&userResponse.Role, &userResponse.CreatedAt, &userResponse.UpdatedAt)
	if err != nil {
		return nil, err
	}
	
	return &userResponse, nil
}


	// Method to update an existing user in the database.
	// It modifies the user details based on the provided ID and returns the updated user response.
	// parameters:
	// 		ctx: context for request-scoped values and cancellation.
	// 		id: the unique identifier of the user to be updated.
	// 		user: the user models containing the updated details.
	// returns:
	// 		UserResponse: a pointer to UserResponse containing the updated user details.
	//  	error: if any occurred during the operation.
func (r *UserRepository) UpdateUser(ctx context.Context, id string, user *models.UpdateUserRequest) (*models.UserResponse, error) {
	query := r.SQKBuilder.
		Update("users").
		Set("user_name", user.UserName).
		Set("email", user.Email).
		Set("password", user.Password).
		Where(sq.Eq{"id": id}).
		Suffix("RETURNING id, user_name, email, role, created_at, updated_at")

	row := query.RunWith(r.db).QueryRowContext(ctx)

	var resp models.UserResponse
	err := row.Scan(&resp.ID, &resp.UserName, &resp.Email,&resp.Role, &resp.CreatedAt, &resp.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found, return nil
		}
		return nil, err
	}
	
	return &resp, nil
}

	// Method to delete a user from the database.
	// It removes the user with the specified ID from the "users" table.
	// parameters:
	// 		ctx: context for request-scoped values and cancellation.
	// 		id: the unique identifier of the user to be deleted.
	// returns:
	// 		error: an error if any occurred during the operation.
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	query := r.SQKBuilder.Delete("users").Where(sq.Eq{"id": id})
	result, err := query.RunWith(r.db).ExecContext(ctx)
	if err != nil {
		return err
	}

	// Get affected rows count
	affected, _ := result.RowsAffected()
	
	
	if affected == 0 {
		return sql.ErrNoRows // No user found to delete
	}
	
	return nil
}

	// method to authenticate a user by email.
	// It retrieves the user details including password for authentication purposes.
	// parameters:
	// 		ctx: context for request-scoped values and cancellation.
	// 		email: the email of the user to be authenticated.
	// returns:
	// 		UserResponse: a pointer to UserResponse containing user details if found.
	//  	error: an error if any occurred during the operation.
func (r *UserRepository) Authenticate(ctx context.Context, email string) (*models.User, error) {
	query := r.SQKBuilder.
		Select("id, user_name, email, password, role").
		From("e_procurement.users").
		Where(sq.Eq{"email": email}).
		Limit(1)

	row := query.RunWith(r.db).QueryRowContext(ctx)

	var user models.User
	err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

	// Method to check if a user exists by email.
	// It queries the "users" table to see if a user with the specified email exists.
	// parameters:
	// 		ctx: context for request-scoped values and cancellation.
	// 		email: the email of the user to check for existence.
	// returns:
	// 		bool: true if the user exists, false otherwise.
func (r *UserRepository) IsUserExists(ctx context.Context, email string) (bool, error) {

	query := r.SQKBuilder.
		Select("1").
		From("users").
		Where(sq.Eq{"email": email}).
		Limit(1)

	row := query.RunWith(r.db).QueryRowContext(ctx)

	var exists int
	err := row.Scan(&exists)
	if err == sql.ErrNoRows {
		// Log successful query but no user found
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

	// Method to get the total count of users in the database.
	// It queries the "users" table to count the total number of users.
	// parameters:
	// 		ctx: context for request-scoped values and cancellation.
	// returns:
	// 		int64: the total count of users.
	// 		error: error if any occurred during the operation.
func (r *UserRepository) GetTotalCount(ctx context.Context) (int64, error) {


	query := r.SQKBuilder.
		Select("COUNT(*)").
		From("users")

	row := query.RunWith(r.db).QueryRowContext(ctx)

	var count int64
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// updatePassword updates the password of a user in the database.
// It modifies the user's password based on the provided ID and returns the updated user response.
// parameters:
// 		ctx: context for request-scoped values and cancellation.
// 		id: the unique identifier of the user whose password is to be updated.
// 		newPassword: the new password to be set for the user.
// // returns:
// 		errors: an error if any occurred during the operation.
func (r *UserRepository) UpdatePassword(ctx context.Context, id string, newPassword string) error {
	query := r.SQKBuilder.
		Update("users").
		Set("password", newPassword).
		Where(sq.Eq{"id": id})

	result, err := query.RunWith(r.db).ExecContext(ctx)
	if err != nil {
		return err
	}
	// Get affected rows count
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return sql.ErrNoRows // No user found to update
	}

	return nil
}