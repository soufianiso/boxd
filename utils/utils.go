package utils

import (
	"encoding/json"
	"net/http"
)

type ApiHandler func(http.ResponseWriter, *http.Request) error 


type ApiError struct {
	Error string
}


// The MiddlewearApi function is a wrapper of my handlers to handler errors 

func WriteJson(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type","application-json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(v)
	return 
}

func WriteError(w http.ResponseWriter, status int, v any){
	w.Header().Set("Content-Type","application-json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(v)
	return  
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
