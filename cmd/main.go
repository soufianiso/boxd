package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/soufianiso/boxd/cmd/api"
	"github.com/redis/go-redis/v9"
	// "context"
)


func main(){
	logger := log.New(os.Stdout, "app: ", log.LstdFlags)

	godotenv.Load()

	// database
	conn := os.Getenv("postgres")
	db, err := sql.Open("postgres",conn)	
	if err != nil{
		logger.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping() ; err != nil{
		logger.Fatal(err)
	}

	// redis connection
	redisOptions := &redis.Options{
		Addr: os.Getenv("REDIS_URL"), 
	}
	redisClient := redis.NewClient(redisOptions)
	defer redisClient.Close()
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		logger.Fatal("redis connection failed", err)
	}

	log.Printf("redis succecufelly connected")


	app := api.NewServer(logger, db, redisClient) 
	server := &http.Server{
		Addr: ":8000",
		Handler: app,
	}

	log.Printf("listening on %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}

}


