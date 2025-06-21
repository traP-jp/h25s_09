package domain

import (
	"time"

	"github.com/google/uuid"
)

type MessageReaction struct {
	MessageID uuid.UUID
	UserName  uuid.UUID
	CreatedAt time.Time
}

type UserReactionToMessage struct {
	MessageID uuid.UUID
	Username  uuid.UUID
	CreatedAt time.Time
}
