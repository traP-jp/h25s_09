package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_09/domain"
)

func (h *handler) ReactionsGetter(c echo.Context) error {

	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid ID")
	}

	//IDが存在しているか
	_, err = h.repo.GetMessageByID(ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "ID is not founded"})
	}
	//リアクションの数の取得
	var count int
	s, _ := h.repo.GetReactionsToMessage(ID)
	count = len(s)

	//ユーザーネームの取得
	usernameRaw := c.Get("username")
	username, ok := usernameRaw.(string)
	if !ok || username == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	myreaction := contains(s, username)

	res := map[string]interface{}{
		"count":      count,
		"myReaction": myreaction,
	}

	return c.JSON(http.StatusCreated, res)

}
func contains(slice []domain.MessageReaction, target string) bool {
	for _, v := range slice {
		if v.Username == target {
			return true
		}
	}
	return false
}
