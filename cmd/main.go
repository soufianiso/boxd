package main

import (
	"database/sql"
	"log"
	"os"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/soufianiso/boxd/cmd/api"
)


func main(){
	godotenv.Load()
	conn := os.Getenv("postgres")
	dbchannel := make(chan *sql.DB)
	fmt.Println("main function start ")	

	go func (){
		newdb, err := sql.Open("postgres",conn)	
		if err != nil{
			log.Fatal(err)
		}
		dbchannel <- newdb 

		
	}()

	db := <- dbchannel
	defer db.Close()

	if err := db.Ping() ; err != nil{
		log.Fatal(err)
	}

	log.Printf("connection success")
	
	app := api.NewServer(":8000",db) 
	app.Run()


	

}

