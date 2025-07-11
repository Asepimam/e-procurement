package https

import (
	"e-procurement/internals/domain/models"
	"e-procurement/internals/usecases"
	response "e-procurement/pkg/responses"
	"e-procurement/pkg/validator"
	"encoding/json"
	"net/http"
)

type AuthHttp struct {
	usecase usecases.AuthUseCase
	validator *validator.CustomValidator
}
func NewAuthHttp(u usecases.AuthUseCase) *AuthHttp {
	return &AuthHttp{
		usecase: u,
		validator: validator.Getvalidator(),
	}
}

func (h *AuthHttp) Authentication(w http.ResponseWriter, r *http.Request){
	// get email and password on body
	var req models.LoginRequest
	

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	
	if err:= h.validator.Validate(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	user,token, err := h.usecase.Authenticate(r.Context(),&req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	response.Success(w, "authenticated successfully", user,nil,token)
}

func (h *AuthHttp) Create(w http.ResponseWriter, r *http.Request) {
	var userReq models.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validator.Validate(userReq); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.usecase.Create(r.Context(),&userReq)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "User created successfully", result, nil)
}