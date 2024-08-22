package user

import (
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
	// "encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soufianiso/boxd/utils"

	// "github.com/soufianiso/boxd/types"
	"time"
)

// "database/sql"




type Handler struct{
	storage Store
}

func NewHandler(storage Store) *Handler {
	return &Handler{ storage : storage }
}



func(h *Handler) SetRoutes(r *mux.Router) *mux.Router{

	r.HandleFunc("/user", utils.MiddlewearApi(h.handleLogin)).Methods("POST")
	return r

}

func(h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) error{
	// Define the signing key 
	signingKey := []byte("secret")	
	
	
	// Create the claims
	claims := jwt.MapClaims{
		"username": "soufiane",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	
	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your secret key
	tokenString, err := token.SignedString(signingKey)
	if err != nil{
		fmt.Println("Error signing key")
		return err
	}

	// user := types.User{}
	// err = json.NewDecoder(r.Body).Decode(&user)
	// if err != nil{
	// 	return err
	// }

	utils.WriteJson(w, http.StatusOK , map[string]string{"autherization":tokenString}) 

	return nil

}








