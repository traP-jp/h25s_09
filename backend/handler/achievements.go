package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type achievement struct {
	Name       string    `json:"name"`
	AchievedAt time.Time `json:"achievedAt"`
}

func (h *handler) PostAchievementsHandler(ctx echo.Context) error {
	username := ctx.Get("username").(string)

	var reqBody struct {
		Name string `json:"name"`
	}
	if err := ctx.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	if reqBody.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	domainAchievement, err := h.repo.InsertUserAchievement(username, reqBody.Name)
	if err != nil {
		ctx.Logger().Error("Failed to insert achievement:", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	achievement := achievement{
		Name:       domainAchievement.AchievementName,
		AchievedAt: domainAchievement.AchievedAt,
	}
	return ctx.JSON(http.StatusCreated, achievement)
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
			Name:       a.AchievementName,
			AchievedAt: a.AchievedAt,
		}
	}
	return ctx.JSON(http.StatusOK, result)
}
