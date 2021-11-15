package routes

import (
	"kelasd/constant"
	"kelasd/controllers/news"
	"kelasd/controllers/users"
	"kelasd/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	middlewares.LogMiddleware(e)

	eNews := e.Group("/news")
	// eNews.Use(middleware.BasicAuth(middlewares.BasicAuth))
	config := middleware.JWTConfig{
		Claims:     &middlewares.JwtCustomClaims{},
		SigningKey: []byte(constant.SECRET_JWT),
	}
	eNews.Use(middleware.JWTWithConfig(config))

	eNews.POST("/", news.InsertNewsController)
	eNews.GET("/", news.GetNewsController)
	eNews.GET("/:newsId", news.GetDetailNewsController)

	e.POST("/login", users.LoginController)

	return e
}
