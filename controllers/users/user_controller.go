package users

import (
	"kelasd/models/response"
	"kelasd/models/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	// email := c.FormValue("email")
	// password := c.FormValue("password")

	var user users.User
	c.Bind(&user)

	if user.Password == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			http.StatusBadRequest,
			"Password Kosong",
			nil,
		})
	}

	// login ke db

	return c.JSON(http.StatusOK, response.BaseResponse{
		http.StatusOK,
		"succes",
		user,
	})

}
