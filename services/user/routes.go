package user

import (
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/soufianiso/boxd/types"

	"os"

	"github.com/joho/godotenv"

	// "encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soufianiso/boxd/auth"
	"github.com/soufianiso/boxd/utils"
)

func SetRoutes(
	r *mux.Router, 
	storage Store, 
	logger *log.Logger, 
	redisClient *redis.Client,
)	{
	r.Handle("/login", handleLogin(storage ,logger)).Methods("POST")
	r.Handle("/register", handleRegister(storage, logger)).Methods("POST")
	}

func handleLogin(storage Store, logger *log.Logger) http.Handler{
	godotenv.Load()
	jwtsecret := os.Getenv("jwtsecret")
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := new(types.User)
		if err := json.NewDecoder(r.Body).Decode(&user) ; err != nil {
			logger.Println("failed decoding",err)
			return 
		}

		u , err := storage.GetUserByEmail(user.Email)
		if err !=  nil{
			logger.Println(err)
			utils.WriteError(w,http.StatusBadRequest, utils.ApiError{ Error: "email or password is incorrect" })
			return 
		}


		if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
			logger.Println(err)
			utils.WriteError(w,http.StatusBadRequest, utils.ApiError{ Error: "email or password is incorrect" })
			return 
		}
		
		tokenString, err := auth.Createjwt(user.Email, jwtsecret)
		if err != nil{
			logger.Println(err)
			return 
		}

		w.Header().Set("Authorization", "Bearer "+tokenString)	
		utils.WriteJson(w, http.StatusOK , map[string]string{"Authorization": tokenString}) 

	})
}

func handleRegister(storage Store, logger *log.Logger) http.Handler{
	user := new(types.User)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		if err := json.NewDecoder(r.Body).Decode(&user) ; err != nil{
			logger.Println(err)
			return 
		}
		
		// Checking whether the email exists or not
		_ , err := storage.GetUserByEmail(user.Email)
		if err ==  nil{
			logger.Println("email already exists")
			utils.WriteError(w,http.StatusBadRequest, utils.ApiError{ Error: "email already exists"})
			return
		}

		// hash the password
		hashedPassword, err := auth.HashPassword(user.Password)
		if err != nil{
			logger.Println(err)
			return 
		}
		
		if err := storage.CreateUser(user, hashedPassword) ; err != nil {
			logger.Println(err)
			return
		}
		
		utils.WriteJson(w, http.StatusCreated, map[string]string{"status":"created"})

	})
}
