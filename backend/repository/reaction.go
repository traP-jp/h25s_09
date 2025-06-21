package repository

import (
	"github.com/google/uuid"
	"github.com/traP-jp/h25s_09/domain"
)

type MessageRactionRepository interface {
	GetMessageReaction(message_id uuid.UUID) (domain.MessageReaction, error)
	InsertMessageReaction(message_id uuid.UUID) (domain.MessageReaction, error)
	DeleteMessageReaction(message_id uuid.UUID) (domain.MessageReaction, error)
}
