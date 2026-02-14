package products

import (
	"log"
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
	err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	products := struct {
		Products []string `json:"products"`
	}{}

	json.Write(w, http.StatusOK, products)
}