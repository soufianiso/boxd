package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/soufianiso/boxd/services/user"
	"github.com/soufianiso/boxd/utils"
)

func NewServer(db *sql.DB) http.Handler{
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	// user service	
	userstore := user.NewStorage(db)
	userHandler := user.UserHandler(userstore) 
	userHandler.SetRoutes(apiRouter)


	// here is top level middleware stuff
	handler := utils.CORSMiddleware(router)
	return handler
}




