package movies

import (
	"net/http"

	"github.com/gorilla/mux"
)



type Handler struct{
	store Store
}



func MoviesHandler(store Store) *Handler{
	return &Handler{store: store}
}



func (h *Handler) SetRoutes( r *mux.Router) *mux.Router{
	r.HandleFunc("/movies", h.listMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", h.listMovieById).Methods("GET")
	
	return  r
}

func (h *Handler) listMovies(w http.ResponseWriter, r *http.Request ){

	
	return  
}


func (h *Handler) listMovieById(w http.ResponseWriter, r *http.Request ){

	

	return 
}







