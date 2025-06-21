package handler

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_09/domain"
	"github.com/traP-jp/h25s_09/handler/middleware"
)

func (h *handler) PostMessageHandler(c echo.Context) error {
	author, ok := c.Get(middleware.UsernameKey).(string)
	if !ok || author == "" {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	message := c.FormValue("message")
	if message == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Message is empty")
	}

	parentIDString := c.FormValue("parent_id")
	parentID := uuid.Nil
	if parentIDString != "" {
		var err error
		parentID, err = uuid.Parse(c.FormValue("parent_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid parent ID")
		}
	}

	file, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid image file")
	}
	// Check MIME type: image/*
	if !strings.HasPrefix(file.Header.Get("Content-Type"), "image/") {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid image file type")
	}
	// Check file size: max 16MiB
	if file.Size > 16*1024*1024 {
		return echo.NewHTTPError(http.StatusBadRequest, "Image file is too large. (max: 16MiB)")
	}

	if parentID != uuid.Nil {
		parent, err := h.repo.GetMessageByID(parentID)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, "Parent message not found")
			}
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve parent message")
		}
		if parent.ParentID != uuid.Nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Cannot reply to a reply")
		}
	}

	fileReader, err := file.Open()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to open image file")
	}
	imageData, err := io.ReadAll(fileReader)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to read image file")
	}

	msg, err := h.repo.CreateMessage(author, message, parentID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create message")
	}
	_, err = h.repo.CreateMessageImage(msg.ID, imageData, file.Header.Get("Content-Type"))
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save image")
	}

	return c.JSON(http.StatusOK, nil)
}
