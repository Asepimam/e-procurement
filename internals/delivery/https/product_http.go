package https

import (
	"e-procurement/internals/domain/models"
	"e-procurement/internals/usecases"
	response "e-procurement/pkg/responses"
	"encoding/json"
	"net/http"
)

type ProductHttp struct {
	productusecase usecases.ProductUseCase
}

func NewProductHttp(u usecases.ProductUseCase) *ProductHttp {
	return&ProductHttp{
		productusecase: u,
	}
}	

// method for http create new product vendor
func(h *ProductHttp) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// decode request body to product model
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