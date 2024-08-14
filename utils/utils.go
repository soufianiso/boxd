package utils

import (
	"log"
	"net/http"
)

type ApiHandler func(http.ResponseWriter, *http.Request) error 


// The MiddlewearApi function is a wrapper of my handlers to handler errors in one function
func MiddlewearApi(f ApiHandler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		err := f(w,r)  
		if err != nil{
			log.Panic(err)
		}
	}
}


func WriteHeaders(w http.ResponseWriter) error {
	w.Header().Set("Content-Type","application-json")
	w.WriteHeader(http.StatusOK)
	return nil
}





