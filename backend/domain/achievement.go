package domain

import (
	"github.com/google/uuid"
	"time"
)

type UserAchievement struct {
	AchievementID	int64
	UserID			uuid.UUID
	AchievedAt      time.Time
}