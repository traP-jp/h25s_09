package repository

import (
	"h25s_09/backend/domain"

	"github.com/google/uuid"
)

type MessageRepository interface {
	CreateMessage(message *domain.Message) (*domain.Message, error)
	GetMessageByID(id uuid.UUID) (*domain.Message, error)
}
