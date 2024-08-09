package main

import (
	"github.com/labstack/echo"
	"github.com/soufiansio/letterboxd/handler"
)



func main(){
	app := echo.New()
	userHandler := handler.UserHandler()
	app.Start(":8080")
}
