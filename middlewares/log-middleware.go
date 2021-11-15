package middlewares

import (
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func LogMiddleware(e *echo.Echo) {
	e.Use(mid.LoggerWithConfig(mid.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}`,
	}))
}
