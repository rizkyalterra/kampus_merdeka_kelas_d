package routes

import (
	"kelasd/controllers/news"
	"kelasd/controllers/users"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// routing
	e.POST("/news", news.InsertNewsController)
	e.GET("/news", news.GetNewsController)
	e.GET("/news/:newsId", news.GetDetailNewsController)
	e.POST("/login", users.LoginController)

	return e
}
