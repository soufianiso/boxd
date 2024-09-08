package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "os/signal"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/soufianiso/boxd/cmd/api"
	// "context"
)


func main(){
	logger := log.New(os.Stdout, "app: ", log.LstdFlags)

	godotenv.Load()

	conn := os.Getenv("postgres")

	db, err := sql.Open("postgres",conn)	
	if err != nil{
		logger.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping() ; err != nil{
		logger.Fatal(err)
	}


	app := api.NewServer(logger,db) 

	server := &http.Server{
		Addr: ":8000",
		Handler: app,
	}

	go func ()  {
		log.Printf("listening on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
		
	} ()

	// listen for system signals for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	logger.Println("Shutting down the server...")

	// Create a deadline to wait for the server to shut down
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %s\n", err)
	}
}


