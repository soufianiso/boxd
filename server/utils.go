package server


import (
	"net/http"
)

type ApiHandler func(http.ResponseWriter, *http.Request) error 



// The MiddlewearApi function is a wrapper of my handlers to handler errors in one function
func MiddlewearApi(f ApiHandler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		err := f(w,r)  
		if err != nil{
		}
	}
}

