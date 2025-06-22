package handler

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_09/domain"
)

func (h *handler) ReactionsDeleter(c echo.Context) error {
	//idを取得してuuid型に
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid ID")
	}
	// idのメッセージがそもそも存在するか
	_, err = h.repo.GetMessageByID(ID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "id not found")
		} //404用
		c.Logger().Error("Failed to delete:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete:")
	} //404以外は500に
	//ユーザーネームの取得
	username := c.Get("username").(string)
	//リアクションを削除
	err = h.repo.DeleteMessageReaction(ID, username)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "id not found")
		} //404用
		c.Logger().Error("Failed to delete:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete")
	} //404以外は500に
	//リアクションの数の取得
	s, err := h.repo.GetReactionsToMessage(ID)
	if err != nil {
		c.Logger().Error("Failed to get reactions:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get reactions")
	}
	count := len(s)
	//自身のリアクションの有無
	myreaction := containsb(s, username)
	res := map[string]interface{}{
		"count":      count,
		"myReaction": myreaction,
	}
	return c.JSON(http.StatusOK, res)
}
func containsb(slice []*domain.MessageReaction, target string) bool {
	for _, v := range slice {
		if v.Username == target {
			return true
		}
	}
	return false
}
