package routes

import (
	userController "keld/controllers/users"

	"github.com/labstack/echo/v4"
)

type RouteControllerList struct {
	UserController userController.UserController
}

func (controller RouteControllerList) RouteRegister(c *echo.Echo) {
	users := c.Group("/user")
	users.POST("/login", controller.UserController.Login)

}
