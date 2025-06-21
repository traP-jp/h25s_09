package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/traP-jp/h25s_09/repository"
)

func Start() {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())

	_, err := repository.NewDB()
	if err != nil {
		e.Logger.Fatal("Failed to connect to the database:", err)
	}

	g := e.Group("/api")
	{
		g.GET("/health", func(c echo.Context) error {
			return c.JSON(200, map[string]string{"status": "ok"})
		})
	}

	e.Logger.Fatal(e.Start(":8080"))
}
