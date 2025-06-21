package domain

import (
	"time"

	"github.com/google/uuid"
)

type MessageImage struct {
	ID        uuid.UUID
	MessageID uuid.UUID
	Data      []byte
	Mime      string
	CreatedAt time.Time
}

type PostMessageImageRequest struct {
	MessageID uuid.UUID
	Data      []byte
	Mime      string
}
