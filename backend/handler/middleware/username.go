package middleware

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

const UsernameKey = "username"
const DevelopmentEnv = "development"

func UsernameProvider(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Request().Header.Get("X-Forwarded-User")
		if username != "" {
			c.Set(UsernameKey, username)
		} else if os.Getenv("ENVIRONMENT") == DevelopmentEnv {
			c.Set(UsernameKey, "anonymous")
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return next(c)
	}
}
