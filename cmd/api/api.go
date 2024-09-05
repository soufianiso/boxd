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



type Api struct {
	port string
	db	*sql.DB 
}


func NewServer(port string, db *sql.DB) *Api{
	return &Api{ 
		port: port, 
		db : db,
	}

}

func(a Api) Run(){
	router := mux.NewRouter()

	userstore := user.NewStorage(a.db)
	userHandler := user.NewHandler(userstore) 
	userHandler.SetRoutes(router)

	moviesStore := movies.NewStorage(a.db)
	moviesHandler := movies.NewHandler(moviesStore) 
	moviesHandler.SetRoutes(router)

	handler := utils.CORSMiddleware(router)
	//here top level http stuff
	s := http.Server{
		Addr: a.port,	
		Handler: handler,
	}

	s.ListenAndServe()
}



