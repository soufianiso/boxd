package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"

	"github.com/soufianiso/boxd/services/user"
	"github.com/soufianiso/boxd/services/movies"
	"github.com/soufianiso/boxd/utils"
)

func NewServer(logger *log.Logger, db *sql.DB, redisClient *redis.Client ) http.Handler{
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStorage(db)
	user.SetRoutes(apiRouter, userStore, logger, redisClient)

	moviesStore := movies.NewStorage(db)
	movies.SetRoutes(apiRouter, moviesStore, logger, redisClient)

	var handler http.Handler
	handler = utils.CORSMiddleware(router)
	// here is top level middleware stuff
	return handler
}




