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

type ProductHttp struct {
	productusecase usecases.ProductUseCase
	validator *validator.CustomValidator
}

func NewProductHttp(u usecases.ProductUseCase) *ProductHttp {
	return&ProductHttp{
		productusecase: u,
		validator: validator.Getvalidator(),
	}
}	

// method for http create new product vendor
func(h *ProductHttp) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productReq models.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&productReq); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// call usecase to create product
	productResponse, err := h.productusecase.CreateProducUsecase(r.Context(), &productReq)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "Product created successfully", productResponse, nil)
}

// method for http get all products
func(h *ProductHttp) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	limitstr, pagestr := r.URL.Query().Get("limit"), r.URL.Query().Get("page")

	limit, err := strconv.Atoi(limitstr)
	if err != nil || limit <= 0 {
		limit = 10 // default limit
	}
	page, err := strconv.Atoi(pagestr)
	if err != nil || page <= 0 {
		page = 1 // default page
	}
	products, count, err := h.productusecase.GetAllProducts(r.Context(), limit, page)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	meta := &response.Meta{
		Page: page, 	
		PerPage: page,
		HasMore: (page*limit) < count,
	}
	response.Success(w, "Products retrieved successfully", products,meta,)
}

// method for http get products by category
func(h *ProductHttp) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		response.Error(w, http.StatusBadRequest, "Category is required")
		return
	}
	category := parts[len(parts)-1]
	if category == "" {
		response.Error(w, http.StatusBadRequest, "Category is required")
		return
	}
	// validate category
	if !h.validator.IsValidUUID(category) {
		response.Error(w, http.StatusBadRequest, "Invalid category format")
		return
	}
	limitstr, pagestr := r.URL.Query().Get("limit"), r.URL.Query().Get("page")
	limit, err := strconv.Atoi(limitstr)
	if err != nil || limit <= 0 {
		limit = 10 // default limit
	}
	page, err := strconv.Atoi(pagestr)
	if err != nil || page <= 0 {
		page = 1 // default page
	}


	products, count, err := h.productusecase.GetProductsByCategory(r.Context(), category, limit, page)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	meta := &response.Meta{
		Page:    page,
		PerPage: limit,
		HasMore: (page*limit) < count,
	}
	response.Success(w, "Products by category retrieved successfully", products, meta)
}

// method for http get product by id
func(h *ProductHttp) GetProductByID(w http.ResponseWriter, r *http.Request) {
	// get product id from url path
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		response.Error(w, http.StatusBadRequest, "Product ID is required")
		return
	}
	id := parts[len(parts)-1] // get the last part of the path as ID
	// validate product id is not empty
	if id == "" {
		response.Error(w, http.StatusBadRequest, "Product ID is required")
		return
	}
	if !h.validator.IsValidUUID(id) {
		response.Error(w, http.StatusBadRequest, "Invalid product ID format")
		return
	}

	product, err := h.productusecase.GetProductByID(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}


	response.Success(w, "Product retrieved successfully", product, nil)
}

// method for http update product
func(h *ProductHttp) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.UpdateProductRequest
	// get product id from path
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		response.Error(w, http.StatusBadRequest, "Product ID is required")
		return
	}
	id := parts[len(parts)-1] // get the last part of the path as ID
	if id == "" {
		response.Error(w, http.StatusBadRequest, "Product ID is required")
		return
	}
	// validate uuid format
	if !h.validator.IsValidUUID(id) {
		response.Error(w, http.StatusBadRequest, "Invalid product ID format")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if h.validator.Validate(product) != nil {
		response.Error(w, http.StatusBadRequest, "Invalid product data")
		return
	}

	resp, err := h.productusecase.UpdateProduct(r.Context(), id, &product)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "Product updated successfully", resp, nil)
}

// delete product by id
func(h *ProductHttp) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// get product id from url path
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		response.Error(w, http.StatusBadRequest, "Product ID is required")
		return
	}
	id := parts[len(parts)-1] // get the last part of the path as ID
	if id == "" {
		response.Error(w, http.StatusBadRequest, "Product ID is required")
		return
	}
	if !h.validator.IsValidUUID(id) {
		response.Error(w, http.StatusBadRequest, "Invalid product ID format")
		return
	}

	err := h.productusecase.DeleteProduct(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "Product deleted successfully", nil, nil)
}