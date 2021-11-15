package users

import (
	"kelasd/configs"
	"kelasd/middlewares"
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

	if err := configs.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			http.StatusBadRequest,
			"Email dan Password tidak sesuai",
			nil,
		})
	}

	if user.Password == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			http.StatusBadRequest,
			"Password Kosong",
			nil,
		})
	}

	// generate token
	token, err := middlewares.GenereteTokenJWT(int(user.Id))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			http.StatusInternalServerError,
			"Error ketika generate token JWT",
			nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		http.StatusOK,
		"succes",
		map[string]string{
			"token": token,
		},
	})

}
