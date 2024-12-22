package api

import (
	"basic_go_backend/services/product"
	user "basic_go_backend/services/user"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// user handler
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	// product handler
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(userStore, productStore)
	productHandler.RegisterProductRoutes(subrouter)

	log.Println("Listening on port", s.addr)
	return http.ListenAndServe(s.addr, router)
}
