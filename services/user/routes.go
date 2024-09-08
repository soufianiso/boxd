package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/soufianiso/boxd/auth"
	"github.com/soufianiso/boxd/types"
	"github.com/soufianiso/boxd/utils"
)

// "database/sql"




type Handler struct{
	logger *log.Logger
	storage Store
}

func UserHandler(logger *log.Logger, storage Store) *Handler {
	return &Handler{ 
		storage : storage, 
		logger : logger, 


	}
}

func(h *Handler) SetRoutes(r *mux.Router) *mux.Router{
	r.HandleFunc("/login", utils.ErrorHandler(h.handleLogin)).Methods("POST")
	r.HandleFunc("/register", utils.ErrorHandler(h.handleRegister)).Methods("POST")
	r.Handle("/test", handleTest(h.logger,h.storage)).Methods("POST")
	return r
}

func(h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) error{
	user := new(types.User)

	// Convert request body to object
	if err := json.NewDecoder(r.Body).Decode(&user) ; err != nil {
		return err
	}

	_ , err := h.storage.GetUserByEmail(user.Email)
	if err !=  nil{
		utils.WriteError(w,http.StatusBadRequest, utils.ApiError{ 
			Error: "email or password incorrect",
		})

		return err
	}

	//create Signed the jwt token and create it 
	godotenv.Load()
	jwtsecret := os.Getenv("jwtsecret")


	tokenString, err := auth.Createjwt(user.Email, jwtsecret)
	if err != nil{
		return err
	}

	return utils.WriteJson(w, http.StatusOK , map[string]string{"Authorization": tokenString}) 
}

func(h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) error{
	user := new(types.User)
	// Convert request body to object
	if err := json.NewDecoder(r.Body).Decode(&user) ; err != nil{
		return err
	}
	
	// Checking whether the email exists or not
	_ , err := h.storage.GetUserByEmail(user.Email)
	if err ==  nil{
		return utils.WriteError(w,http.StatusBadRequest, utils.ApiError{ 
			Error: "email or password incorrect",
		})
	}
	
	// hash the password
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil{
		return err
	}
	
	if err := h.storage.CreateUser(user, hashedPassword) ; err != nil {
		return err
	}

	return utils.WriteJson(w, http.StatusCreated, map[string]string{"status":"created"})
}



func handleTest(logger *log.Logger, storage Store) http.Handler{
	logger.Print("test")
	s := "test"
	return http.HandlerFunc( 
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w,s)		
			},
		)
	
}
