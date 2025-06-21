package handler

import (
	"net/http"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_09/domain"
)

func (h *handler) GetMessageByMessageID(ctx echo.Context) error {
	// パラメータからmessageIdを取得してUUIDに変換
	messageIDstr := ctx.Param("messageId")
	if messageIDstr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "messageId is required")
	}
	messageIDuuid, err := uuid.Parse(messageIDstr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid messageId format")
	}
	// メッセージを取得
	message, err := h.repo.GetMessageByID(messageIDuuid)
	if err != nil {
		if err == domain.ErrNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Message not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve message")
	}
	return ctx.JSON(http.StatusOK, message)
}
