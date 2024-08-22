package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type ApiHandler func(http.ResponseWriter, *http.Request) error 


type ApiError struct {
	Error string
}


// The MiddlewearApi function is a wrapper of my handlers to handler errors 
func MiddlewearApi(f ApiHandler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		err := f(w,r)  
		if err != nil{
			log.Panic(err)
			WriteJson(w, http.StatusBadRequest, ApiError{Error: "invalid request"})
		}
	}
}


func WriteHeaders(w http.ResponseWriter) error {
	w.Header().Set("Content-Type","application-json")
	w.WriteHeader(http.StatusOK)
	return nil
}




func WriteJson(w http.ResponseWriter, status int, v any) error{
	w.Header().Set("Content-Type","application-json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)

}
