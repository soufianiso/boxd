package user

import (
	"log"
	"github.com/redis/go-redis/v9"
	"github.com/soufianiso/boxd/types"
	"os"
	"github.com/joho/godotenv"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/soufianiso/boxd/auth"
	"github.com/soufianiso/boxd/utils"
)
 

// Routes
func SetRoutes( r *mux.Router, storage Store, logger *log.Logger, redisClient *redis.Client,) {
	r.Handle("/login", handleLogin(storage ,logger)).Methods("POST")
	r.Handle("/register", handleRegister(storage, logger)).Methods("POST")
	}

// Handlers
func handleLogin(storage Store, logger *log.Logger) http.Handler{
	godotenv.Load()
	jwtsecret := os.Getenv("jwtsecret")
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user types.User 

		if err := utils.Decode(r, &user) ; err != nil{
			logger.Printf("Failed to decode request body:%v",err)
			return
		}

		if err := utils.Validate(&user); err != nil{
			logger.Printf("validation failed: %s", err)
			utils.Encode(w, r, http.StatusBadRequest, utils.ApiError{ Error: "email or password is incorrect" })
			return 
		}

		u , err := storage.GetUserByEmail(user.Email)
		if err !=  nil{
			utils.Encode(w, r, http.StatusBadRequest, utils.ApiError{ Error: "email or password is incorrect" })
			logger.Printf("Failed to retrieve user by Email (%s)",user.Email)
			return 
		}

		if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
			logger.Printf("Password mismatch for the user: (%s)",user.Email)
			utils.Encode(w, r, http.StatusBadRequest, utils.ApiError{ Error: "email or password is incorrect" })
			return 
		}
		
		tokenString, err := auth.Createjwt(user.Email, jwtsecret)
		if err != nil{
			logger.Printf("failed to create JWT token: %v",err)
			return 
		}

		err = utils.Encode(w, r, http.StatusCreated, map[string]string{"Authorization": tokenString}) 
		if err != nil{
			logger.Printf("failed to encode the response: %v",err)
		}
	})
}

func handleRegister(storage Store, logger *log.Logger) http.Handler{
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user types.User

		if err := utils.Decode(r, &user) ; err != nil{
			logger.Printf("Failed to decode the body: %v", err)
			return
		}

		// Checking whether the email exists or not
		exists , err := storage.GetUserByEmail(user.Email)
		if exists !=  nil{
			logger.Printf("Failed to retrieve user by email (%s)", user.Email)
			utils.Encode(w, r,  http.StatusBadRequest, utils.ApiError{ Error: "email already exists"})
			return
		}

		// hash the password
		hashedPassword, err := auth.HashPassword(user.Password)
		if err != nil{
			logger.Printf("failed to hash password %s",user.Password)
			return 
		}
		
		if err := storage.CreateUser(&user, hashedPassword) ; err != nil {
			logger.Printf("failed to create user %s",user.Email)
			return
		}
		
		err = utils.Encode(w, r, http.StatusCreated, map[string]string{"status":"created"}) 
		if err != nil{
			logger.Printf("failed to encode the response: %v",err)
		}

	})
}
