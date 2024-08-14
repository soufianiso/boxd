package main

import (
	"github.com/soufianiso/letterboxd/server"
)


func main(){
	app := server.NewServer(":8080") 
	app.Run()

}
