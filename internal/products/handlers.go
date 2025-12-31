package products

import (
	"log"
	"net/http"

	"github.com/wydentis/ecommerce/internal/json"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ListProduct(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Printf("products handler handle: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.WriteJSON(w, http.StatusOK, products)
}
