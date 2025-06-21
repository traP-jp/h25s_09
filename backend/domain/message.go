package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID         uuid.UUID
	Author     string
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}