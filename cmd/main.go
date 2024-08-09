package main

import (
	"github.com/soufianiso/letterboxd/handler"
	"github.com/labstack/echo"
)



func main(){
	app := echo.New()
	userHandler := handler.UserHandler()
	app.Start(":8080")
}
