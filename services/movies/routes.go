package movies

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func SetRoutes(r *mux.Router, storage Store, logger *log.Logger, redisClient *redis.Client){
	r.Handle("/movies", handleListMovies(storage ,logger, redisClient)).Methods("GET")
	r.Handle("/movies/{id}", handleGetMovie(storage ,logger ,redisClient)).Methods("GET")
}

func handleListMovies(storage Store, logger *log.Logger, redisClient *redis.Client) http.Handler{
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
	})
}

func handleGetMovie(storage Store, logger *log.Logger, redisClient *redis.Client) http.Handler{
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
