package main

import (
	"log"
	"os"
	// "database/sql"
	// _ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"github.com/soufianiso/letterboxd/server"
	"github.com/soufianiso/letterboxd/storage"
)


func main(){


	godotenv.Load()
	conn := os.Getenv("postgres")

	store, err := storage.NewStore(conn)
	if err != nil {
		log.Println("the bot", conn)
	}

	app := server.NewServer(":8000",store) 
	app.Run()

	// psql_uri := "user=postgres  password=password dbname=postgres port=5432 sslmode=disable"

	// db, err := sql.Open("postgres",psql_uri) 
	// if err != nil{
	// 	panic(err)
	// }
	// _ , err = db.Exec(`CREATE TABLE users ( username VARCHAR(50), password TEXT)`)	
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := db.Ping(); err != nil{
	// 	log.Fatal(err)		
	// }

	

}
