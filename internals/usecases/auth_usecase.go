package usecases

import (
	"context"
	"e-procurement/internals/domain/models"
	"e-procurement/internals/repositories"
	"e-procurement/pkg/auth"
	"e-procurement/pkg/encripted"
	"errors"
	"fmt"
	"log"
)

type AuthUseCase struct {
	repo *repositories.UserRepository
	encripted *encripted.Encripted
	jwt *auth.JWT
}

func NewAuthUseCase(
	repo *repositories.UserRepository,
	JWT *auth.JWT,
	) *AuthUseCase {
	return &AuthUseCase{
		repo: repo,
		jwt: JWT,
		encripted: encripted.NewEncripted(),
	}
}

func(u *AuthUseCase) Authenticate(ctx context.Context, data *models.LoginRequest)(*models.UserResponse,string, error) {
	
	log.Printf("Authenticating user with email: %s", data.Email)
	user, err := u.repo.Authenticate(ctx, data.Email)
	log.Printf("User found: %v", user)
	if err != nil {
		log.Printf("Error authenticating user: %v", err)
		return nil,"", errors.New("invalid email or password")
	}

	if user == nil {
		return nil,"", errors.New("invalid email or password")
	}

	isValid, err := u.encripted.CheckPasswordHash(user.Password, data.Password)
	if err != nil {
		return nil,"", errors.New("error checking password")
	}

	if !isValid {
		return nil,"",errors.New("invalid email or password")
	}

	token, err := u.jwt.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, "", errors.New("error generating token")
	}
	userResponse := &models.UserResponse{
		ID:        	user.ID,
		UserName: 	user.UserName,
		Email:     	user.Email,
		Role: 		user.Role,
		CreatedAt: 	user.CreatedAt,
		UpdatedAt: 	user.UpdatedAt,
	}
	return userResponse,token, nil
}


func (u *AuthUseCase) Create(ctx context.Context, user *models.CreateUserRequest) (*models.UserResponse, error) {
	exists, err := u.repo.IsUserExists(ctx, user.Email)
	if err != nil {
		return nil, fmt.Errorf("error checking user existence: %v", err)
	}
	if exists {
		return nil, errors.New("user already exists")
	}

	user.Password, err = u.encripted.HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %v", err)
	}

	resp, err := u.repo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}
	userResponse := &models.UserResponse{
		ID:        resp.ID,
		UserName:  resp.UserName,
		Email:     resp.Email,
		Role:      resp.Role,
		CreatedAt: resp.CreatedAt,
		UpdatedAt: resp.UpdatedAt,
	}

	return userResponse, nil
}
