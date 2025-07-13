package https

import (
	"e-procurement/internals/domain/models"
	"e-procurement/internals/usecases"
	response "e-procurement/pkg/responses"
	"e-procurement/pkg/validator"
	"encoding/json"
	"net/http"
)

type UserHttp struct {
	userUseCase usecases.UserUseCase
	validator *validator.CustomValidator
}

func NewUserHttp(u usecases.UserUseCase) *UserHttp {
	return &UserHttp{
		userUseCase: u,
		validator: validator.Getvalidator(),
	}
}

// GetUserByID handles the HTTP request to get user details by ID.
func (h *UserHttp) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		response.Error(w, http.StatusBadRequest, "user ID is required")
		return
	}
	if !h.validator.IsValidUUID(userID) {
		response.Error(w, http.StatusBadRequest, "invalid user ID format")
		return
	}												

	// Call usecase to get user by ID
	userResponse, err := h.userUseCase.GetUserByID(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "User details retrieved successfully", userResponse, nil)
}

func (h *UserHttp) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userReq models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validate the request
	if err := h.validator.Validate(userReq); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Call usecase to update user
	userResponse, err := h.userUseCase.UpdateUser(r.Context(), &userReq)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "User updated successfully", userResponse, nil)
}

func (h *UserHttp) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		response.Error(w, http.StatusBadRequest, "user ID is required")
		return
	}
	if !h.validator.IsValidUUID(userID) {
		response.Error(w, http.StatusBadRequest, "invalid user ID format")
		return
	}

	// Call usecase to delete user
	if err := h.userUseCase.DeleteUser(r.Context(), userID); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "User deleted successfully", nil, nil)
}

func(h *UserHttp) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var changePasswordReq models.ChangePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&changePasswordReq); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validate the request
	if err := h.validator.Validate(changePasswordReq); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Call usecase to change password
	if err := h.userUseCase.UpdatePassword(r.Context(), &changePasswordReq); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "Password changed successfully", nil, nil)
}