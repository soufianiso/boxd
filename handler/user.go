package handler

import (
	"github.com/labstack/echo"
	"github.com/soufianiso/letterboxd/view/user"
	"github.com/soufianiso/letterboxd/model"
)




func HandleUserShow(c echo.Context) error{
	u := model.User{
		Email: "soufiane@gmail.com" ,
	}
	return user.Show(u).Render(c.Request().Context(), c.Response())
		
} 
