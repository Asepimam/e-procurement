package https

import (
	"e-procurement/internals/domain/models"
	"e-procurement/internals/usecases"
	constants "e-procurement/pkg/constans"
	response "e-procurement/pkg/responses"
	"e-procurement/pkg/validator"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type VendorHttp struct {
	vendorusecase usecases.VendorUseCase
	validator      *validator.CustomValidator
}

func NewVendortHttp(u usecases.VendorUseCase) *VendorHttp {
	return &VendorHttp{
		vendorusecase: u,
		validator:      validator.Getvalidator(),
	}
}


func (h *VendorHttp) CreateVendor(w http.ResponseWriter, r *http.Request) {
	// decode request body to vendor model
	var vendorReq models.CreateVendorRequest
	if err := json.NewDecoder(r.Body).Decode(&vendorReq); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	// validate user ID is provided
	userID := r.Context().Value(constants.ContextUserIDKey).( string )
	if !h.validator.IsValidUUID(userID) {
		response.Error(w, http.StatusBadRequest, "Invalid user")
		return
	}
	// validate vendor request
	if err := h.validator.Validate(vendorReq); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// call usecase to create vendor
	vendorResponse, err := h.vendorusecase.CreateVendorUsecase(r.Context(), &vendorReq)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "Vendor created successfully", vendorResponse, nil)
}

func (h *VendorHttp) GetAllVendors(w http.ResponseWriter, r *http.Request) {
	limitstr, pagestr := r.URL.Query().Get("limit"), r.URL.Query().Get("page")
	limit, err := strconv.Atoi(limitstr)
	if err != nil || limit <= 0 {
		limit = 10 // default limit
	}
	page, err := strconv.Atoi(pagestr)
	if err != nil || page <= 0 {
		page = 1 // default page
	}


	vendors, count, err := h.vendorusecase.GetAllVendors(r.Context(), limit, page)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	meta := &response.Meta{
		Page:    page,
		PerPage: limit,
		HasMore: (page * limit) < count,
	}

	response.Success(w, "Vendors retrieved successfully", vendors, meta)
}

func (h *VendorHttp) GetVendorByID(w http.ResponseWriter, r *http.Request) {
	// get vendor id from url path
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		response.Error(w, http.StatusBadRequest, "Vendor ID is required")
		return
	}
	id := parts[len(parts)-1] // get the last part of the path as ID

	// validate vendor id is not empty
	if id == "" {
		response.Error(w, http.StatusBadRequest, "Vendor ID is required")
		return
	}
	if !h.validator.IsValidUUID(id) {
		response.Error(w, http.StatusBadRequest, "Invalid vendor ID format")
		return
	}

	vendorResponse, err := h.vendorusecase.GetVendorByID(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "Vendor retrieved successfully", vendorResponse, nil)
}

func (h *VendorHttp) UpdateVendor(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	vendorID := parts[len(parts)-1] // get the last part of the path as ID
	var vendorReq models.UpdateVendorRequest
	if err := json.NewDecoder(r.Body).Decode(&vendorReq); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	
	if vendorID == "" {
		response.Error(w, http.StatusBadRequest, "Vendor ID is required")
		return
	}

	if !h.validator.IsValidUUID(vendorID) {
		response.Error(w, http.StatusBadRequest, "Invalid vendor ID format")
		return
	}

	if err := h.validator.Validate(vendorReq); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	vendorResponse, err := h.vendorusecase.UpdateVendor(r.Context(), vendorID, &vendorReq)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "Vendor updated successfully", vendorResponse, nil)
}

func(h *VendorHttp) DeleteVendor(w http.ResponseWriter, r *http.Request) {
	// get vendor id from url path
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		response.Error(w, http.StatusBadRequest, "Vendor ID is required")
		return
	}
	id := parts[len(parts)-1] // get the last part of the path as ID

	// validate vendor id is not empty
	if id == "" {
		response.Error(w, http.StatusBadRequest, "Vendor ID is required")
		return
	}
	if !h.validator.IsValidUUID(id) {
		response.Error(w, http.StatusBadRequest, "Invalid vendor ID format")
		return
	}

	err := h.vendorusecase.DeleteVendor(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, "Vendor deleted successfully", nil, nil)
}