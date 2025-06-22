package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type achievement struct {
	Name      string    `json:"name"`
	AchievedAt time.Time `json:"achievedAt"`
}

func (h *handler) PostAchievementsHandler(ctx echo.Context) error {
	username := ctx.Get("username").(string)
	if username == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

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
		Name:      domainAchievement.AchievementName,
		AchievedAt: domainAchievement.AchievedAt,
	}
	return ctx.JSON(http.StatusCreated, achievement)
}