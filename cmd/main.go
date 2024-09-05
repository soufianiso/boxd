package main

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	"net/http"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/soufianiso/boxd/cmd/api"
)


func main(){
	godotenv.Load()

	conn := os.Getenv("postgres")

	db, err := sql.Open("postgres",conn)	
	if err != nil{
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping() ; err != nil{
		log.Fatal(err)
	}

	app := api.NewServer(db) 

	server := &http.Server{
		Addr: ":8000",
		Handler: app,
	}

	go func ()  {
		log.Printf("listening on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
		
	}()
}

