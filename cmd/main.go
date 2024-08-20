package main

import (
	"database/sql"
	"log"
	"os"

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

	if err := db.Ping() ; err != nil{
		log.Fatal(err)
	}

	app := api.NewServer(":8000",db) 
	app.Run()


	

}
