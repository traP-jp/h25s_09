package domain

import (
	"github.com/google/uuid"
	"time"
)

type UserAchievement struct {
	AchievementID	int64  `json:"id"`
	UUID			uuid.UUID `json:"uuid"`
	AchievedAt      	time.Time `json:"achieved_at"`
}