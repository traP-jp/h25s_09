package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetHealthHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
