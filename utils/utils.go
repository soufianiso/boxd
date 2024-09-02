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
func ErrorHandler(f ApiHandler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		err := f(w,r)  
		if err != nil{
			log.Print(err)
			WriteError(w, http.StatusInternalServerError, ApiError{Error: "bad request"})
			return 
		}
	}
}

func WriteJson(w http.ResponseWriter, status int, v any) error{
	w.Header().Set("Content-Type","application-json")
	w.WriteHeader(status)

	err :=  json.NewEncoder(w).Encode(v)
	return err
}

func WriteError(w http.ResponseWriter, status int, v any) error{
	w.Header().Set("Content-Type","application-json")
	w.WriteHeader(status)

	err :=  json.NewEncoder(w).Encode(v)
	return err
}




