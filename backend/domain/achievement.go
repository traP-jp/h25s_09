package domain

import (
	"time"
)

type UserAchievement struct {
	AchievementID int64
	Username      string
	AchievedAt    time.Time
}
