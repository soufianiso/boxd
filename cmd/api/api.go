package api

import (
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/soufianiso/boxd/services/user"
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

	

	http.ListenAndServe(a.port, router)

	return 
}
