package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJson(r *http.Request, payload any) error {
	if payload == nil {
		return fmt.Errorf("missing request body")
	}
	// log.Print("payload", payload, "r.body", json.NewDecoder(r.Body))
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJson(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, map[string]string{"error": err.Error()})
}
func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	token := strings.TrimPrefix(tokenAuth, "Bearer ")

	tokenQuery := r.URL.Query().Get("token")
	if tokenAuth != "" {
		return token
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}
