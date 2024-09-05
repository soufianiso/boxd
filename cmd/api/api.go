package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"

	// "github.com/gorilla/handlers"
	"github.com/soufianiso/boxd/services/movies"
	"github.com/soufianiso/boxd/services/user"
	"github.com/soufianiso/boxd/utils"
)

func NewServer(db *sql.DB) http.Handler{
	router := mux.NewRouter()

	// user service	
	userstore := user.NewStorage(db)
	userHandler := user.UserHandler(userstore) 
	userHandler.SetRoutes(router)

	// movies service	
	moviesStore := movies.NewStorage(db)
	moviesHandler := movies.MoviesHandler(moviesStore) 
	moviesHandler.SetRoutes(router)

	// here is top level middleware stuff
	handler := utils.CORSMiddleware(router)
	return handler
}



