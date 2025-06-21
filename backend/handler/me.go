package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	m "github.com/traP-jp/h25s_09/handler/middleware"
)

func (h *handler) GetMeHandler(ctx echo.Context) error {
	result := ctx.Get(m.UsernameKey).(string)
	if result == "" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}
	return ctx.JSON(http.StatusOK, map[string]string{"traqId": result})
}
