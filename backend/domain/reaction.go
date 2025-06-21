package domain

import (
	"time"

	"github.com/google/uuid"
)

type MessageReaction struct {
	Message_id uuid.UUID
	Username   uuid.UUID
	Created_at time.Time
}
