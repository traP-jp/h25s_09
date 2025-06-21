package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())

	g := e.Group("/api")
	{
		g.GET("/health", func(c echo.Context) error {
			return c.JSON(200, map[string]string{"status": "ok"})
		})
	}

	e.Logger.Fatal(e.Start(":8080"))
}
