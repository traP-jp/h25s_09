package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	m "github.com/traP-jp/h25s_09/handler/middleware"
	"github.com/traP-jp/h25s_09/repository"
)

type handler struct {
	repo repository.Repository
}

func Start() {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())
	e.Use(m.UsernameProvider)

	db, err := repository.NewDB()
	if err != nil {
		e.Logger.Fatal("Failed to connect to the database:", err)
	}
	h := &handler{
		repo: repository.NewRepository(db),
	}

	g := e.Group("/api")
	{
		g.GET("/health", h.GetHealthHandler)
		g.GET("/images/:id", h.GetMessageImageHandler)
		g.GET("/try-achieve/:id", h.TryAchieveHandler)
		u := g.Group("/users/:name")
		{
			u.GET("/achievements", h.GetUserAchievementsHandler)
		}
		me := g.Group("/me")
		{
			me.GET("", h.GetMeHandler)
			me.GET("/achievements", h.GetMyAchievementsHandler)
		}
		msg := g.Group("/messages")
		{
			msg.GET("", h.GetMessagesHandler)
			msg.POST("", h.PostMessageHandler)      
		  msg.GET("/:id", h.GetMessageHandler)
			msg.POST("/:id/reactions", h.ReactionsAdder)
			msg.DELETE("/:id/reactions", h.ReactionsDeleter)
		}
	}

	e.Logger.Fatal(e.Start(":8080"))
}
