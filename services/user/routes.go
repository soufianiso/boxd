package user

import (
	"context"
	"encoding/json"

	"github.com/soufianiso/boxd/types"

	"os"

	"github.com/joho/godotenv"

	"log"
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

	// Convert request body to object
	if err := json.NewDecoder(r.Body).Decode(&user) ; err != nil {
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
    ctx := r.Context()
	user := new(types.User)
	// Convert request body to object
	if err := json.NewDecoder(r.Body).Decode(&user) ; err != nil{
		return err
	}
	
	// Checking whether the email exists or not
	_ , err := h.storage.GetUserByEmail(ctx,user.Email)
	if err == context.Canceled{
		// Client canceled the request, so we stop processing
		log.Println("Request canceled by the client")
		return nil
	}

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

	select {
    case <-ctx.Done(): // Handle request cancellation
        if ctx.Err() == context.Canceled {
            log.Println("Request was canceled by the client")
            return nil
        }
    }
	return utils.WriteJson(w, http.StatusCreated, map[string]string{"status":"created"})
}






