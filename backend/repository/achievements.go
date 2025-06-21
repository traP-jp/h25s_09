package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/traP-jp/h25s_09/domain"
)

type AchievementsRepository interface {
	GetUserAchievements(userID uuid.UUID) ([]domain.UserAchievement, error)
	InsertUserAchievement(userID uuid.UUID, achievementID int64) (*domain.UserAchievement, error)
}

type achievementsRepositoryImpl struct {
	db *sqlx.DB
}

func NewAchievementsRepository(db *sqlx.DB) AchievementsRepository {
	return &achievementsRepositoryImpl{db: db}
}

func (r *achievementsRepositoryImpl) GetUserAchievements(userID uuid.UUID) ([]domain.UserAchievement, error) {
	return nil, domain.ErrNotImplemented
}

func (r *achievementsRepositoryImpl) InsertUserAchievement(userID uuid.UUID, achievementID int64) (*domain.UserAchievement, error) {
	return nil, domain.ErrNotImplemented
}
