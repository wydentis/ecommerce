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
	err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Printf("products handler handle: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products := struct {
		Products []string `json:"products"`
	}{[]string{"Agatha", "is", "my", "girlfriend"}}

	json.WriteJSON(w, http.StatusOK, products)
}
