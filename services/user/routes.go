package user

import (
	"encoding/json"

	"github.com/soufianiso/boxd/types"

	// "encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soufianiso/boxd/auth"
	"github.com/soufianiso/boxd/utils"
)

// "database/sql"




type Handler struct{
	storage Store
}

func NewHandler(storage Store) *Handler {
	return &Handler{ storage : storage }
}



func(h *Handler) SetRoutes(r *mux.Router) *mux.Router{

	r.HandleFunc("/login", utils.ErrorHandler(h.handleLogin)).Methods("POST")
	r.HandleFunc("/register", utils.ErrorHandler(h.handleRegister)).Methods("POST")
	return r

}

func(h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) error{
	user := new(types.User)

	if err := json.NewDecoder(r.Body).Decode(&user) ; err != nil {
		return err
	}

	secretkey := "secret"
	tokenString, err := auth.Createjwt(user.Username, secretkey)
	if err != nil{
		return err
	}

	return utils.WriteJson(w, http.StatusOK , map[string]string{"Authorization": tokenString}) 

}

func(h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) error{
	user := new(types.User)
	if err := json.NewDecoder(r.Body).Decode(user) ; err != nil{
		return err
	}
	
	_ , err := h.storage.GetUserByEmail(user.Username)
	if err ==  nil{
		return utils.WriteError(w,http.StatusBadRequest, utils.ApiError{ 
			Error: "email or password incorrect",
		})
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil{
		return err
	}

	if err := h.storage.CreateUser(user.Username, hashedPassword) ; err != nil {
		return err
	}
	
	
	return utils.WriteJson(w, http.StatusCreated, map[string]string{"status":"created"})

}






