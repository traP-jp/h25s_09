package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

var achievements = map[int64]int64{
	1: 1,
	2: 2,
	3: 3,
	4: 4,
	5: 5,
}

func (h *handler) TryAchieveHandler(ctx echo.Context) error {
	username, ok := ctx.Get("username").(string)
	if !ok || username == "" {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	ID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid ID")
	}

	_, exists := achievements[ID]
	if !exists {
		return echo.NewHTTPError(http.StatusNotFound, "achievement not found")
	}

	// ここに何かしらの判定ロジックを追加
	Judge := true
	if !Judge {
		return ctx.JSON(http.StatusOK, map[string]bool{"dispatched": false})
	}

	hasAchievement, err := h.hasUserAchievement(username, ID)
	if err != nil {
		ctx.Logger().Error("Failed to check user achievements:", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if !hasAchievement {
		_, err = h.repo.InsertUserAchievement(username, ID)
		if err != nil {
			ctx.Logger().Error("Failed to insert user achievement:", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	return ctx.JSON(http.StatusOK, map[string]bool{"dispatched": true})
}

func (h *handler) hasUserAchievement(username string, achievementID int64) (bool, error) {
	userAchievements, err := h.repo.GetUserAchievements(username)
	if err != nil {
		return false, err
	}

	for _, achievement := range userAchievements {
		if achievement.AchievementID == achievementID {
			return true, nil
		}
	}
	return false, nil
}

type achievement struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	AchievedAt time.Time `json:"achievedAt"`
}

func (h *handler) GetUserAchievementsHandler(ctx echo.Context) error {
	username := ctx.Param("name")
	if username == "" {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	userAchievements, err := h.repo.GetUserAchievements(username)
	if err != nil {
		ctx.Logger().Error("Failed to retrieve user achievements:", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	result := make([]*achievement, len(userAchievements))
	for i, a := range userAchievements {
		result[i] = &achievement{
			ID:         a.AchievementID,
			Name:       "",
			AchievedAt: a.AchievedAt,
		}
	}
	return ctx.JSON(http.StatusOK, result)
}
