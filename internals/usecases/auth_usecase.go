package usecases

import (
	"context"
	"e-procurement/internals/domain/models"
	"e-procurement/internals/repositories"
	"errors"
	"fmt"
)

type AuthUseCase struct {
	repo *repositories.UserRepository

}

func NewAuthUseCase(
	repo *repositories.UserRepository,
	) *AuthUseCase {
	return &AuthUseCase{
		repo: repo,
	}
}

func(u *AuthUseCase) Authenticate(ctx context.Context, data *models.LoginRequest)(*models.UserResponse, error) {
	user, err := u.repo.Authenticate(ctx, data.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// isValid, err := u.encript.CheckPasswordHash(user.Password, data.Password)
	// if err != nil {
	// 	logerContext.Warn("error checking password hash", zap.String("email", data.Email), zap.Error(err))
	// 	return nil, "", errors.New("error checking password")
	// }

	// if !isValid {
	// 	logerContext.Warn("invalid password attempt", zap.String("email", data.Email))
	// 	return nil, "", errors.New("invalid email or password")
	// }

	// token, err := u.jwt.GenerateToken(user.ID, profile.Position)
	// if err != nil {
	// 	logerContext.Error("error generating token", zap.String("user_id", user.ID), zap.Error(err))
	// 	return nil, "", errors.New("error generating token")
	// }
	userResponse := &models.UserResponse{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return userResponse, nil
}


func (u *AuthUseCase) Create(ctx context.Context, user *models.CreateUserRequest) (*models.UserResponse, error) {
	exists, err := u.repo.IsUserExists(ctx, user.Email)
	if err != nil {
		return nil, fmt.Errorf("error checking user existence: %v", err)
	}
	if exists {
		return nil, errors.New("user already exists")
	}

	// user.Password, err = u.encript.HashPassword(user.Password)
	// if err != nil {
	// 	return nil, fmt.Errorf("error hashing password: %v", err)
	// }

	resp, err := u.repo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	return resp, nil
}
