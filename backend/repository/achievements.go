package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/traP-jp/h25s_09/domain"
)

type AchievementsRepository interface {
	GetUserAchievements(username string) ([]domain.UserAchievement, error)
	InsertUserAchievement(username string, achievementName string) (*domain.UserAchievement, error)
}

type userAchievement struct {
	AchievementName string    `db:"name"`
	Username        string    `db:"username"`
	AchievedAt      time.Time `db:"achieved_at"`
}

func (r *repositoryImpl) GetUserAchievements(username string) ([]domain.UserAchievement, error) {
	var achievementsDB []userAchievement
	err := r.db.Select(&achievementsDB, "SELECT name, username, achieved_at FROM achievements WHERE username = ?", username)
	if err != nil {
		return nil, err
	}

	// ドメインの型に変換
	domainAchievements := make([]domain.UserAchievement, 0, len(achievementsDB))
	for _, achDB := range achievementsDB {
		domainAchievements = append(domainAchievements, domain.UserAchievement{
			AchievementName: achDB.AchievementName,
			Username:        achDB.Username,
			AchievedAt:      achDB.AchievedAt,
		})
	}

	return domainAchievements, nil
}

func (r *repositoryImpl) InsertUserAchievement(username string, achievementName string) (*domain.UserAchievement, error) {
	_, err := r.db.Exec("INSERT INTO achievements (name, username) VALUES (?, ?)", achievementName, username)
	if err != nil {
		return nil, err
	}
	var achievement userAchievement
	err = r.db.Get(&achievement, "SELECT name, username, achieved_at FROM achievements WHERE name = ? AND username = ?", achievementName, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	// ドメインの型に変換
	domainAchievement := domain.UserAchievement{
		AchievementName: achievement.AchievementName,
		Username:        achievement.Username,
		AchievedAt:      achievement.AchievedAt,
	}

	return &domainAchievement, nil
}
