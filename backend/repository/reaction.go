package repository

import (
	"github.com/google/uuid"
	"github.com/traP-jp/h25s_09/domain"
)

type MessageRactionRepository interface {
	GetMessageReaction(messageId uuid.UUID) (*domain.MessageReaction, error)
	InsertMessageReaction(messageId uuid.UUID) (*domain.MessageReaction, error)
	DeleteMessageReaction(messageId uuid.UUID) (*domain.MessageReaction, error)
}
