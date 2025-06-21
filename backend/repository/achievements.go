package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_09/domain"
)

type AchievementsRepository interface {
	GetUserAchievements(userID uuid.UUID) ([]domain.UserAchievement, error)
	InsertUserAchievement(userID uuid.UUID, achievementID int64) (*domain.UserAchievement, error)
}

type userAchievement struct {
	AchievementID int64     `db:"id"`
	Username      uuid.UUID `db:"username"`
	AchievedAt    time.Time  `db:"achieved_at"`
}

func (r *repositoryImpl) GetUserAchievements(userID uuid.UUID) ([]domain.UserAchievement, error) {
	var achievementsDB []userAchievement
	err := r.db.Select(&achievementsDB, "SELECT * FROM achievements WHERE username = ?", userID)
	if err != nil {
		return nil, err
	}

	// ドメインの型に変換
	domainAchievements := make([]domain.UserAchievement, 0, len(achievementsDB))
	for _, achDB := range achievementsDB {
		domainAchievements = append(domainAchievements, domain.UserAchievement{
			AchievementID: achDB.AchievementID,
			UserID:        achDB.Username,
			AchievedAt:    achDB.AchievedAt,
		})
	}

	return domainAchievements, nil
}

func (r *repositoryImpl) InsertUserAchievement(userID uuid.UUID, achievementID int64) (*domain.UserAchievement, error) {
	_, err := r.db.Exec("INSERT INTO achievements (id, username) VALUES (?, ?)", achievementID, userID)
	if err != nil {
		return nil, err
	}
	var achievement userAchievement
	err = r.db.Get(&achievement, "SELECT * FROM achievements WHERE id = ? AND username = ?", achievementID, userID)
	if err != nil {
		return nil, err
	}

	// ドメインの型に変換
	domainAchievement := domain.UserAchievement{
		AchievementID: achievement.AchievementID,
		UserID:        achievement.Username,
		AchievedAt:    achievement.AchievedAt,
	}

	return &domainAchievement, nil
}