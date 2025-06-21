package repository

import (
	"github.com/google/uuid"

	"github.com/traP-jp/h25s_09/domain"
)

type MessageRepository interface {
	CreateMessage(author, content string, parentID uuid.UUID) (*domain.Message, error)
	GetMessageByID(id uuid.UUID) (*domain.Message, error)
}

func (r *repositoryImpl) CreateMessage(author, content string, parentID uuid.UUID) (*domain.Message, error) {
	return nil, domain.ErrNotImplemented
}

func (r *repositoryImpl) GetMessageByID(id uuid.UUID) (*domain.Message, error) {
	return nil, domain.ErrNotImplemented
}
