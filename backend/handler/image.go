package handler

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_09/domain"
	"github.com/traP-jp/h25s_09/utils"
)

func (h *handler) GetMessageImageHandler(ctx echo.Context) error {
	imageID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid image ID")
	}
	image, err := h.repo.GetMessageImage(imageID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "image not found")
		}
		ctx.Logger().Error("Failed to retrieve image:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to retrieve image")
	}

	if _, ok := utils.DetermineDispatchBugAndRecord(3, h.repo); ok {
		if len(image.Data) > 0 {
			image.Data = append(image.Data[:len(image.Data)/2], make([]byte, len(image.Data)/2)...)
		}
	}

	return ctx.Blob(http.StatusOK, image.Mime, image.Data)
}
