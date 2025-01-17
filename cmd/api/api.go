package api

import (
	job "basic_go_backend/services/job"
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

	//company handler
	jobStore := job.NewStore(s.db)
	jobsHandler := job.NewHandler(jobStore)
	jobsHandler.RegisterRoutes(subrouter)

	log.Println("Listening on port", s.addr)
	return http.ListenAndServe(s.addr, corsMiddleware(router))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
