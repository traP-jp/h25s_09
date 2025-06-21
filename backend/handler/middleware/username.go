package middleware

import (
	"os"

	"github.com/labstack/echo/v4"
)

const UsernameKey = "username"

func UsernameProvider(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Request().Header.Get("X-Forwarded-User")
		if username != "" {
			c.Set(UsernameKey, username)
		} else if os.Getenv("ENVIRONMENT") == "development" {
			c.Set(UsernameKey, "anonymous")
		}
		return next(c)
	}
}
