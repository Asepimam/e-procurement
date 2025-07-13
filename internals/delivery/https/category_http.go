package https

import (
	"e-procurement/internals/domain/models"
	"e-procurement/internals/usecases"
	response "e-procurement/pkg/responses"
	"e-procurement/pkg/validator"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type CategoryHttp struct {
	categoryUsecase usecases.CategoryUsecase
	validator *validator.CustomValidator

}

func NewCategoryHttp(u usecases.CategoryUsecase) *CategoryHttp {
	return &CategoryHttp{
		categoryUsecase: u,
		validator: validator.Getvalidator(),
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

	// validate category request
	if err := h.validator.Validate(categoryReq); err != nil {
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

func (h *CategoryHttp) GetCategory(w http.ResponseWriter, r *http.Request) {
	limitstr, pagestr := r.URL.Query().Get("limit"), r.URL.Query().Get("page")
	limit, err := strconv.Atoi(limitstr)
	if err != nil || limit <= 0 {
		limit = 10 // default limit
	}
	page, err := strconv.Atoi(pagestr)
	if err != nil || page <= 0 {
		page = 1 // default page
	}

	categories,count, err := h.categoryUsecase.GetAllCategoriesUsecase(r.Context(), limit, page)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	meta := &response.Meta{
		Page:    page,
		PerPage: limit,
		HasMore: (page*limit) < count,
	}
	response.Success(w, "Get Category", categories,meta)
}

func (h *CategoryHttp) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	// gunakan url path untuk mendapatkan ID kategori
	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		response.Error(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	id := parts[len(parts)-1] // get the last part of the path as ID
	// validasi ID is not empty
	if id == "" {
		response.Error(w, http.StatusBadRequest, "ID is required")
		return
	}

	valid := h.validator.IsValidUUID(id)
	if !valid {
		response.Error(w, http.StatusBadRequest, "Invalid category ID format")
		return
	}
	category, err := h.categoryUsecase.GetCategoryByIDUsecase(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "Get Category By ID", category, nil)
}

func (h *CategoryHttp) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// decode request body to update category model
	var categoryReq models.UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&categoryReq); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// validate category request
	if err := h.validator.Validate(categoryReq); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// gunakan url path untuk mendapatkan ID kategori
	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		response.Error(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	id := parts[len(parts)-1] // get the last part of the path as ID
	if id == "" {
		response.Error(w, http.StatusBadRequest, "ID is required")
		return
	}

	valid := h.validator.IsValidUUID(id)
	if !valid {
		response.Error(w, http.StatusBadRequest, "Invalid category ID format")
		return
	}

	categoryResponse, err := h.categoryUsecase.UpdateCategoryUsecase(r.Context(), id, &categoryReq)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "Category updated successfully", categoryResponse, nil)
}

func (h *CategoryHttp) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// get url path to get category ID
	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		response.Error(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	id := parts[len(parts)-1] // get the last part of the path as ID
	if id == "" {
		response.Error(w, http.StatusBadRequest, "ID is required")
		return
	}

	valid := h.validator.IsValidUUID(id)
	if !valid {
		response.Error(w, http.StatusBadRequest, "Invalid category ID format")
		return
	}

	err := h.categoryUsecase.DeleteCategoryUsecase(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "Category deleted successfully", nil, nil)
}