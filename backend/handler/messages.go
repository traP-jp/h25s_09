package handler

import (
	"cmp"
	"errors"
	"io"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_09/domain"
	"github.com/traP-jp/h25s_09/handler/middleware"
)

type reactions struct {
	Count      int64 `json:"count"`
	MyReaction bool  `json:"myReaction"`
}

type message struct {
	ID         uuid.UUID `json:"id"`
	Author     string    `json:"author"`
	Content    string    `json:"content"`
	ImageID    uuid.UUID `json:"imageId,omitempty"`
	Reactions  reactions `json:"reactions"`
	ReplyCount int64     `json:"replyCount"`
	CreatedAt  time.Time `json:"createdAt"`
}

func (h *handler) GetMessagesHandler(ctx echo.Context) error {
	var messages []domain.Message
	var err error

	limit, err := strconv.ParseInt(cmp.Or(ctx.QueryParam("limit"), "20"), 10, 64)
	if err != nil || limit <= 0 || limit > 100 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid limit parameter")
	}
	offset, err := strconv.ParseInt(cmp.Or(ctx.QueryParam("offset"), "0"), 10, 64)
	if err != nil || offset < 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid offset parameter")
	}
	traqID := ctx.QueryParam("traqId")
	includeReplies := ctx.QueryParam("includeReplies") == "true"

	// Fetch messages from the repository
	messages, err = h.repo.GetMessages(limit, offset, traqID, includeReplies)
	if err != nil {
		ctx.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	jsonMessages := make([]message, len(messages))
	for i, msg := range messages {
		ImageID, err := h.repo.GetMessageImageIDByMessageID(msg.ID)
		if err != nil {
			if !errors.Is(err, domain.ErrNotFound) {
				ctx.Logger().Error(err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			ImageID = uuid.Nil
		}
		Replies, err := h.repo.GetRepliesByMessageID(msg.ID)
		if err != nil {
			ctx.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		RepliesCount := int64(len(Replies))
		Reactions, err := h.repo.GetReactionsToMessage(msg.ID)
		if err != nil {
			ctx.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		ReactionsCount := int64(len(Reactions))
		MyReaction := slices.ContainsFunc(Reactions, func(r *domain.MessageReaction) bool {
			return r.Username == ctx.Get("username").(string)
		})

		jsonMessages[i] = message{
			ID:      msg.ID,
			Author:  msg.Author,
			Content: msg.Content,
			ImageID: ImageID,
			Reactions: reactions{
				Count:      ReactionsCount,
				MyReaction: MyReaction,
			},
			ReplyCount: RepliesCount,
			CreatedAt:  msg.CreatedAt,
		}
	}
	return ctx.JSON(http.StatusOK, jsonMessages)
}

const MaxImageSize = 16 * 1024 * 1024 // 16 MiB

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
		parentID, err = uuid.Parse(parentIDString)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid parent ID")
		}
	}

	file, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid image file")
	}
	if file != nil {
		// Check MIME type: image/*
		if !strings.HasPrefix(file.Header.Get("Content-Type"), "image/") {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid image file type")
		}
		// Check file size: max 16MiB
		if file.Size > MaxImageSize {
			return echo.NewHTTPError(http.StatusBadRequest, "Image file is too large. (max: 16MiB)")
		}
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

	var imageData []byte
	if file != nil {
		fileReader, err := file.Open()
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to open image file")
		}
		defer fileReader.Close()
		imageData, err = io.ReadAll(fileReader)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to read image file")
		}
	}

	msg, err := h.repo.CreateMessage(author, message, parentID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create message")
	}
	imgID := uuid.Nil
	if file != nil && len(imageData) != 0 {
		img, err := h.repo.CreateMessageImage(msg.ID, imageData, file.Header.Get("Content-Type"))
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save image")
		}
		imgID = img.ID
	}

	return c.JSON(http.StatusOK, &messageDetail{
		ID:        msg.ID,
		Author:    msg.Author,
		Content:   msg.Content,
		ImageID:   imgID,
		Reactions: reactions{Count: 0, MyReaction: false},
		Replies:   []any{},
		CreatedAt: msg.CreatedAt,
	})
}

type messageDetail struct {
	ID        uuid.UUID `json:"id"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	ImageID   uuid.UUID `json:"imageId,omitempty"`
	Reactions reactions `json:"reactions"`
	Replies   []any     `json:"replies"`
	CreatedAt time.Time `json:"createdAt"`
}
