package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	m "github.com/traP-jp/h25s_09/handler/middleware"
)

func (h *handler) GetMeHandler(ctx echo.Context) error {
	result, ok := ctx.Get(m.UsernameKey).(string)
	if !ok || result == "" {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	return ctx.JSON(http.StatusOK, map[string]string{"traqId": result})
}

func (h *handler) GetMyAchievementsHandler(ctx echo.Context) error {
	username := ctx.Get(m.UsernameKey).(string)
	achievements, err := h.repo.GetUserAchievements(username)
	if err != nil {
		ctx.Logger().Error("Failed to retrieve user achievements:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to retrieve achievements")
	}
	result := make([]achievement, len(achievements))
	for i, a := range achievements {
		result[i] = achievement{
			Name:       a.AchievementName,
			AchievedAt: a.AchievedAt,
		}
	}
	return ctx.JSON(http.StatusOK, result)
}
