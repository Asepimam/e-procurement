package https

import (
	"e-procurement/internals/domain/models"
	"e-procurement/internals/usecases"
	response "e-procurement/pkg/responses"
	"encoding/json"
	"net/http"
)

type CategoryHttp struct {
	categoryUsecase usecases.CategoryUsecase
}

func NewCategoryHttp(u usecases.CategoryUsecase) *CategoryHttp {
	return &CategoryHttp{
		categoryUsecase: u,
	}
}

// method for http create new category
func (h *CategoryHttp) CreateCategory(w http.ResponseWriter, r *http.Request) {
	// decode request body to category model
	var categoryReq models.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&categoryReq); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// call usecase to create category
	categoryResponse, err := h.categoryUsecase.CreateCategoryUsecase(r.Context(), &categoryReq)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "Category created successfully", categoryResponse, nil)
}