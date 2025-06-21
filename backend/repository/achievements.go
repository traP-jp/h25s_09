package repository

import (
	"github.com/google/uuid"
	"github.com/traP-jp/h25s_09/domain"
)

type AchievementsRepository interface {
	GetUserAchievements(userID uuid.UUID) ([]domain.UserAchievement, error)
	InsertUserAchievement(userID uuid.UUID, achievementID int64) (*domain.UserAchievement, error)
}

func (r *repositoryImpl) GetUserAchievements(userID uuid.UUID) ([]domain.UserAchievement, error) {
	return nil, domain.ErrNotImplemented
}

func (r *repositoryImpl) InsertUserAchievement(userID uuid.UUID, achievementID int64) (*domain.UserAchievement, error) {
	return nil, domain.ErrNotImplemented
}
