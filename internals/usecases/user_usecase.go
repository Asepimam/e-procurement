package usecases

import (
	"context"
	"e-procurement/internals/domain/models"
	"e-procurement/internals/repositories"
	customContext "e-procurement/pkg/context"
	"e-procurement/pkg/encripted"
	"fmt"
)

type UserUseCase struct {
	userRepository *repositories.UserRepository
	encripted *encripted.Encripted
}

func NewUserUseCase(userRepo *repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepo,
		encripted: encripted.NewEncripted(),
	}
}

// Method to get user details by ID
func (u *UserUseCase) GetUserByID(ctx context.Context,userID string)(*models.UserResponse, error) {
	user, err := u.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user with ID %s does not exist", userID)
	}

	userResponse := &models.UserResponse{
		ID:        		user.ID,
		UserName:      	user.UserName,
		Email:     		user.Email,
		Role:      		user.Role,
		CreatedAt: 		user.CreatedAt,
		UpdatedAt: 		user.UpdatedAt,
	}

	return userResponse, nil
}

func (u *UserUseCase) UpdateUser(ctx context.Context, userReq *models.UpdateUserRequest) (*models.UpdateUserResponse, error) {
	userID, err := customContext.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user ID from context: %w", err)
	}

	exitsUser, err := u.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	if exitsUser == nil {
		return nil, fmt.Errorf("user with ID %s does not exist", userID)
	}

	if userReq.UserName == "" {
		userReq.UserName = exitsUser.UserName
	}
	if userReq.Email == "" {
		userReq.Email = exitsUser.Email
	}
	if userReq.Password == "" {
		userReq.Password = exitsUser.Password
	}

	user, err := u.userRepository.UpdateUser(ctx, userID, userReq)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	userUpdateResponse := &models.UpdateUserResponse{
		ID:        		user.ID,
		UserName:      	user.UserName,
		Email:     		user.Email,
		CreatedAt: 		user.CreatedAt,
		UpdatedAt: 		user.UpdatedAt,
	}

	return userUpdateResponse, nil
}

// methot to delete user by ID
func (u *UserUseCase) DeleteUser(ctx context.Context,userID string) error {
	existingUser, err := u.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user by ID: %w", err)
	}
	if existingUser == nil {
		return fmt.Errorf("user with ID %s does not exist", userID)
	}

	err = u.userRepository.Delete(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func(u *UserUseCase) UpdatePassword(ctx context.Context, userReq *models.ChangePasswordRequest)  error {
	userID, err := customContext.GetUserIDFromContext(ctx)
	if err != nil {
		return  fmt.Errorf("failed to get user ID from context: %w", err)
	}

	// get existing user
	existingUser, err := u.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user by ID: %w", err)
	}
	if existingUser == nil {
		return fmt.Errorf("user with ID %s does not exist", userID)
	}

	// validate old password
	isValidExistingPassword, err := u.encripted.CheckPasswordHash(existingUser.Password, userReq.OldPassword)
	if err != nil {
		return fmt.Errorf("error checking existing password: %w", err)
	}
	if !isValidExistingPassword {
		return fmt.Errorf("invalid old password")
	}
	if userReq.NewPassword == "" {
		userReq.NewPassword = existingUser.Password
	}
	// hash new password
	userReq.NewPassword, err = u.encripted.HashPassword(userReq.NewPassword)
	if err != nil {
		return fmt.Errorf("error hashing new password: %w", err)
	}

	// proceed to update password
	err = u.userRepository.UpdatePassword(ctx, existingUser.ID, userReq.NewPassword)
	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}