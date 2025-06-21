package handler

import (
	"cmp"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_09/domain"
)

type reactions struct {
	Count    int64 `json:"count"`
	MyReaction   bool  `json:"myReaction"`
}

type message struct {
	ID        uuid.UUID `json:"id"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	ImageID   uuid.UUID `json:"imageId,omitempty"`
	Reactions reactions  `json:"reactions"`
	ReplyCount int64     `json:"replyCount"`
	CreatedAt time.Time  `json:"createdAt"`
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
	traqId := ctx.QueryParam("traqId")
	includeReplies := ctx.QueryParam("includeReplies") == "true"

	// Fetch messages from the repository
	messages, err = h.repo.GetMessages(limit, offset, traqId, includeReplies)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve messages: "+err.Error())
	}

	jsonMessages := make([]message, len(messages))
	for i, msg := range messages {
		ImageID, err := h.repo.GetMessageImageIDByMessageID(msg.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve message image ID: "+err.Error())
		}
		Replies, err := h.repo.GetRepliesByMessageID(msg.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve replies count: "+err.Error())
		}
		RepliesCount := int64(len(Replies))
		Reactions, err := h.repo.GetReactionsToMessage(msg.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve reactions: "+err.Error())
		}
		ReactionsCount := int64(len(Reactions))
		MyReaction := slices.ContainsFunc(Reactions, func(r *domain.MessageReaction) bool {
			return r.Username == ctx.Get("username").(string)
		})


		jsonMessages[i] = message{
			ID:        msg.ID,
			Author:    msg.Author,
			Content:   msg.Content,
			ImageID:   ImageID,
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

