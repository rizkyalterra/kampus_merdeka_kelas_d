package middlewares

import (
	"kelasd/configs"
	"kelasd/models/users"

	"github.com/labstack/echo/v4"
)

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	var db = configs.DB
	var user users.User

	if err := db.Where("email = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil

}
