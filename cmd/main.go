package main

import (
	"github.com/soufianiso/letterboxd/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)


func main(){
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())


	app.GET("/", handler.HandleUserShow)
	app.Start(":8080")
}
