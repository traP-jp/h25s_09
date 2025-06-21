package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_09/domain"
)

func (h *handler) GetMessageImageHandler(ctx echo.Context) error {
	imageID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid image ID")
	}
	image, err := h.repo.GetMessageImage(imageID)
	if err != nil {
		if err == domain.ErrNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "image not found")
		}
		ctx.Logger().Error("Failed to retrieve image:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to retrieve image")
	}
	return ctx.Blob(http.StatusOK, image.Mime, image.Data)
}
