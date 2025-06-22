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
		ctx.Logger().Error("Failed to retrieve messages:", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	jsonMessages := make([]message, len(messages))
	for i, msg := range messages {
		ImageID, err := h.repo.GetMessageImageIDByMessageID(msg.ID)
		if err != nil {
			if !errors.Is(err, domain.ErrNotFound) {
				ctx.Logger().Error("Failed to retrieve image ID for message:", msg.ID, err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			ImageID = uuid.Nil
		}
		Replies, err := h.repo.GetRepliesByMessageID(msg.ID)
		if err != nil {
			ctx.Logger().Error("Failed to retrieve replies for message:", msg.ID, err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		RepliesCount := int64(len(Replies))
		Reactions, err := h.repo.GetReactionsToMessage(msg.ID)
		if err != nil {
			ctx.Logger().Error("Failed to retrieve reactions for message:", msg.ID, err)
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
	author := c.Get(middleware.UsernameKey).(string)

	content := c.FormValue("message")
	if content == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Message is empty")
	}

	parentIDString := c.FormValue("repliesTo")
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

	msg, err := h.repo.CreateMessage(author, content, parentID)
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
		Replies:   []message{},
		CreatedAt: msg.CreatedAt,
	})
}

type messageDetail struct {
	ID        uuid.UUID `json:"id"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	ImageID   uuid.UUID `json:"imageId,omitempty"`
	Reactions reactions `json:"reactions"`
	Replies   []message `json:"replies"`
	CreatedAt time.Time `json:"createdAt"`
}

func (h *handler) GetMessageHandler(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid message ID")
	}

	msg, err := h.repo.GetMessageByID(ID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Message not found")
		}
		c.Logger().Error("Failed to retrieve message:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve message")
	}

	imageID, err := h.repo.GetMessageImageIDByMessageID(ID)
	if err != nil {
		if !errors.Is(err, domain.ErrNotFound) {
			c.Logger().Error("Failed to retrieve image ID for message:", ID, err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve image ID")
		}
		imageID = uuid.Nil
		c.Logger().Info("No image found for message:", ID)
	}

	reactionList, err := h.repo.GetReactionsToMessage(ID)
	if err != nil {
		c.Logger().Error("Failed to retrieve reactions for message:", ID, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve reactions")
	}
	reactionsCount := int64(len(reactionList))
	myReaction := slices.ContainsFunc(reactionList, func(r *domain.MessageReaction) bool {
		return r.Username == c.Get("username").(string)
	})

	replies, err := h.repo.GetRepliesByMessageID(ID)
	if err != nil {
		c.Logger().Error("Failed to retrieve replies for message:", ID, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve replies")
	}

	repliesList := make([]message, len(replies))
	for i, reply := range replies {
		replyImageID, err := h.repo.GetMessageImageIDByMessageID(reply.ID)
		if err != nil {
			if !errors.Is(err, domain.ErrNotFound) {
				c.Logger().Error("Failed to retrieve image ID for reply:", reply.ID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve reply image ID")
			}
			replyImageID = uuid.Nil
		}
		replyReactionList, err := h.repo.GetReactionsToMessage(reply.ID)
		if err != nil {
			c.Logger().Error("Failed to retrieve reactions for reply:", reply.ID, err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve reply reactions")
		}

		repliesList[i] = message{
			ID:      reply.ID,
			Author:  reply.Author,
			Content: reply.Content,
			ImageID: replyImageID,
			Reactions: reactions{
				Count: int64(len(replyReactionList)),
				MyReaction: slices.ContainsFunc(replyReactionList, func(r *domain.MessageReaction) bool {
					return r.Username == c.Get("username").(string)
				}),
			},
			CreatedAt: reply.CreatedAt,
		}
	}

	return c.JSON(http.StatusOK, &messageDetail{
		ID:      ID,
		Author:  msg.Author,
		Content: msg.Content,
		ImageID: imageID,
		Reactions: reactions{
			Count:      reactionsCount,
			MyReaction: myReaction,
		},
		Replies:   repliesList,
		CreatedAt: msg.CreatedAt,
	})
}
