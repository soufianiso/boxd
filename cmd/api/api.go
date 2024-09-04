package api

import (
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
	// "github.com/gorilla/handlers"
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


	http.ListenAndServe(a.port,allowCors(router))
	return 
}

func allowCors(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Replace with your React app's URL
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w,r)
	})
}


