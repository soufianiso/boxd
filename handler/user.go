package handler

import (
	"github.com/labstack/echo"
	"github.com/soufianiso/letterboxd/view/user"
)


type UserHandler struct { }


func (h UserHandler) HandleUserShow(c echo.Context) error{
	return user.Show().Render(c.Request().Context(), c.Response())
		
} 
