package repository

import (
	"github.com/google/uuid"

	"github.com/traP-jp/h25s_09/domain"
)

type MessageRepository interface {
	CreateMessage(message *domain.Message) (*domain.Message, error)
	GetMessageByID(id uuid.UUID) (*domain.Message, error)
}
