package product

import (
	"basic_go_backend/services/auth"
	"basic_go_backend/types"
	"basic_go_backend/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	userStore    types.UserStore
	productStore types.ProductStore
}

func NewHandler(user types.UserStore, store types.ProductStore) *Handler {
	return &Handler{
		userStore:    user,
		productStore: store,
	}
}

func (h *Handler) RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.getProductsHandler).Methods(http.MethodGet)
	router.HandleFunc("/events", h.eventsHandler).Methods(http.MethodGet)
	router.HandleFunc("/products/{productID}", h.getProductHandler).Methods(http.MethodGet)
	router.HandleFunc("/products", auth.WithJWTAuth(h.createProductHandler, h.userStore)).Methods(http.MethodPost)
}

func (h *Handler) getProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.productStore.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, products)
}

func (h *Handler) getProductHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) createProductHandler(w http.ResponseWriter, r *http.Request) {

	var product types.CreateProductPayload
	err := utils.ParseJson(r, &product)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(product); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}
	err = h.productStore.CreateProduct(product)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusCreated, product)
}
func (h *Handler) eventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/event-stream")
	// w.Header().Add("cache-control", "no-cache")
	tokens := []string{"My", "Name", "Is", "Jeno", "And", "I", "Am", "A", "Developer"}
	for _, token := range tokens {
		content := fmt.Sprintf("data: %s\n\n", token)
		w.Write([]byte(content))
		w.(http.Flusher).Flush()
		time.Sleep(1 * time.Second)
	}

}
