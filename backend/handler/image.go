package handler

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
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
	imageObj, err := h.repo.GetMessageImage(imageID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "image not found")
		}
		ctx.Logger().Error("Failed to retrieve image:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to retrieve image")
	}

	if utils.DetermineDispatchBug(ctx, h.repo, h.ss, 3) {
		if len(imageObj.Data) > 0 {
			imageObj.Data = append(imageObj.Data[:len(imageObj.Data)/2], make([]byte, len(imageObj.Data)/2)...)
		}
	} else {
		if _, ok := utils.DetermineDispatchBugAndRecord(8, h.repo); ok {
			// []byte → image.Image に変換
			imgDecoded, _, err := image.Decode(bytes.NewReader(imageObj.Data))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "failed to decode image")
			}
			//image.Image→[]byte
			var buf bytes.Buffer
			err = jpeg.Encode(&buf, imgDecoded, &jpeg.Options{Quality: 5})
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "failed to encode resized image")
			}
			// Dataを書き換え
			imageObj.Data = buf.Bytes()
			imageObj.Mime = "image/jpeg" // MIMEも変更
		}
	}

	return ctx.Blob(http.StatusOK, imageObj.Mime, imageObj.Data)
}
