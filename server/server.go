package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soufianiso/letterboxd/handlers"
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

type Server struct{
	port  string
}


func NewServer(port string) *Server{
	return &Server{port: port }
}


func (s *Server) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/", MiddlewearApi(handlers.GetFilms))
	http.ListenAndServe(s.port, router)
}


