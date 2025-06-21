package domain

import (
	"time"

	"github.com/google/uuid"
)

type MessageReaction struct {
	MessageID uuid.UUID
	Username  string
	CreatedAt time.Time
}
