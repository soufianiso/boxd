package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/soufianiso/boxd/types"
)

type ApiHandler func(http.ResponseWriter, *http.Request) error 


type ApiError struct {
	Error string
}


func CORSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == "OPTIONS" {
            // Respond to preflight request with status 200 OK
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}
func Encode(w http.ResponseWriter, r *http.Request, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return err
	}
	return nil
}

func Decode(r *http.Request, v any) (error) {
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return  fmt.Errorf("decode json: %w", err)
	}
	return  nil
}

func Validate(user *types.User) error {
	switch {
	case len(user.Email) == 0:
		return errors.New("empty form")
	case len(user.Password) == 0:
		return errors.New("empty form")
	default:
		return nil
	}
}
