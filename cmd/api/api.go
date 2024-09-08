package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// "github.com/gorilla/handlers"
	"github.com/soufianiso/boxd/services/movies"
	"github.com/soufianiso/boxd/services/user"
	"github.com/soufianiso/boxd/utils"
)

func NewServer(logger *log.Logger, db *sql.DB) http.Handler{
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	// user service	
	userstore := user.NewStorage(db)
	userHandler := user.UserHandler(logger, userstore) 
	userHandler.SetRoutes(apiRouter)

	// movies service	
	moviesStore := movies.NewStorage(db)
	moviesHandler := movies.MoviesHandler(moviesStore) 
	moviesHandler.SetRoutes(apiRouter)

	// here is top level middleware stuff
	var handler http.Handler
	handler = utils.CORSMiddleware(router)
	return handler
}



