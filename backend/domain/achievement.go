package domain

import (
	"time"
)

type UserAchievement struct {
	AchievementName string
	Username        string
	AchievedAt      time.Time
}
