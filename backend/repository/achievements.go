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
	var achievements []domain.UserAchievement
	err := r.db.Select(&achievements, "SELECT * FROM achievements WHERE username = ?", userID)
	if err != nil {
		return nil, err
	}
	return achievements, nil
}

func (r *repositoryImpl) InsertUserAchievement(userID uuid.UUID, achievementID int64) (*domain.UserAchievement, error) {
	_, err := r.db.Exec("INSERT INTO achievements (id, username) VALUES (?, ?)", achievementID, userID)
	if err != nil {
		return nil, err
	}
	var achievement domain.UserAchievement
	err = r.db.Get(&achievement, "SELECT * FROM achievements WHERE id = ? AND username = ?", achievementID, userID)
	if err != nil {
		return nil, err
	}
	return &achievement, nil
}