package movies

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func SetRoutes(r *mux.Router, storage Store, logger *log.Logger){
	r.Handle("/movies", handleListMovies(storage ,logger)).Methods("GET")
	r.Handle("/register", handleGetMovie(storage, logger)).Methods("GET")
}

func handleListMovies(storage Store, logger *log.Logger) http.Handler{
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
	})
}

func handleGetMovie(storage Store, logger *log.Logger) http.Handler{
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
