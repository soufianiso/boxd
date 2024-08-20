package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// "database/sql"




type Handler struct{
	storage Store
}

func NewHandler(storage Store) *Handler {
	return &Handler{ storage : storage }
}



func(h *Handler) SetRoutes(r *mux.Router) *mux.Router{
	r.HandleFunc("/user", h.userhandle )
	return r

}

func(h *Handler) userhandle(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"ss")

}








