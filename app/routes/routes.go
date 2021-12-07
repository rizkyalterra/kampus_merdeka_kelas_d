package routes

import (
	userController "keld/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	UserController userController.UserController
	JWTConfig      middleware.JWTConfig
}

func (controller RouteControllerList) RouteRegister(c *echo.Echo) {
	users := c.Group("/user")
	// users.Use(middleware.JWTWithConfig(controller.JWTConfig))
	users.POST("/login", controller.UserController.Login)
	users.PUT("/:userId", controller.UserController.Login)
	users.POST("/login", controller.UserController.Login)

	users.POST("/cart/", controller.UserController.Login)

	// users.GET("/", controller.UserController.Login, middleware.JWTWithConfig(controller.JWTConfig))
}
