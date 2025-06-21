package domain

import (
	"time"

	"github.com/google/uuid"
)

type MessageReaction struct {
	MessageID uuid.UUID
	UserName  string
	CreatedAt time.Time
}

type GetMessageReactionResponse struct {
	Count              int
	UserAlreadyReacted bool
}

type InsertMessageReactionResponce struct {
	MessageID uuid.UUID
	UserName  string
	CreatedAt time.Time
	GetMessageReactionResponse 
}
