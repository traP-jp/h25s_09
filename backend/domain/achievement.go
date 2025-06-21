package domain

import "github.com/google/uuid"

type UserAchievement struct {
	AchievementID	int64  `json:"id"`
	UUID			uuid.UUID `json:"uuid"`
	AchivedAt      	string `json:"achived_at"`
}