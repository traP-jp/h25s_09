package domain

import (
	"github.com/google/uuid"
	"time"
)

type UserAchievement struct {
	AchievementID	int64
	Username		string
	AchievedAt      time.Time
}