package domain

import (
	"github.com/google/uuid"
	"time"
)

type UserAchievement struct {
	AchievementID	int64
	UUID			uuid.UUID
	AchievedAt      time.Time
}