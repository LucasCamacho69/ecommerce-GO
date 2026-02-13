package products

import (
	"net/http"

	"github.com/LucasCamacho69/ecommerce-GO/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProduct(w http.ResponseWriter, r *http.Request) {
	
	products := []string{"Hello", "World"}

	json.Write(w, http.StatusOK, products)
}