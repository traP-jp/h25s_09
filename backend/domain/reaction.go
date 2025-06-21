package domain

import (
	"time"

	"github.com/google/uuid"
)

type MessageReaction struct {
	MessageID uuid.UUID
	UserName   string
	CreatedAt time.Time
}
