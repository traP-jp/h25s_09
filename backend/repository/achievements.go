package repository

import (
	"github.com/google/uuid"
	"github.com/traP-jp/h25s_09/domain"
)

type AchievementsRepository interface {
	GetUserAchievements(userId uuid.UUID) ([]domain.UserAchievement, error)
	InsertUserAchievement(userId uuid.UUID, achievementID int64) (domain.UserAchievement, error)
}