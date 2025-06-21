package domain

import (
	"time"

	"github.com/google/uuid"
)

type MessageImage struct {
	Id        uuid.UUID
	MessageId uuid.UUID
	Data      []byte
	Mime      string
	CreatedAt time.Time
}

type PostMessageImageRequest struct {
	MessageId uuid.UUID
	Data      []byte
	Mime      string
}
