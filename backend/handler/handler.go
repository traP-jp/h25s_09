package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/traP-jp/h25s_09/repository"
)

type handler struct {
	repo repository.Repository
}

func Start() {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())

	db, err := repository.NewDB()
	if err != nil {
		e.Logger.Fatal("Failed to connect to the database:", err)
	}
	_ = &handler{
		repo: repository.NewRepository(db),
	}

	g := e.Group("/api")
	{
		g.GET("/health", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
		})
	}

	e.Logger.Fatal(e.Start(":8080"))
}
