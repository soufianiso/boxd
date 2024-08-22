package user

import (
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
	// "encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soufianiso/boxd/utils"
	"github.com/soufianiso/boxd/types"
)

// "database/sql"




type Handler struct{
	storage Store
}

func NewHandler(storage Store) *Handler {
	return &Handler{ storage : storage }
}



func(h *Handler) SetRoutes(r *mux.Router) *mux.Router{
	r.HandleFunc("/user", utils.MiddlewearApi(h.userhandle)).Methods("POST")
	return r

}

func(h *Handler) userhandle(w http.ResponseWriter, r *http.Request) error{

	user := types.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		return err
	}

	utils.WriteJson(w, http.StatusOK , user) 

	return nil

}








