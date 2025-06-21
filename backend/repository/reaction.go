package repository

import (
	"github.com/google/uuid"
	"github.com/traP-jp/h25s_09/domain"
)

type MessageRactionRepository interface {
	GetMessageReaction(messageID uuid.UUID) (*domain.MessageReaction, error)
	InsertMessageReaction(messageID uuid.UUID) (*domain.MessageReaction, error)
	DeleteMessageReaction(messageID uuid.UUID) (*domain.MessageReaction, error)
}

func (r *repositoryImpl) GetMessageReaction(messageID uuid.UUID) (*domain.MessageReaction, error) {
	return nil, domain.ErrNotImplemented
}

func (r *repositoryImpl) InsertMessageReaction(messageID uuid.UUID) (*domain.MessageReaction, error) {
	return nil, domain.ErrNotImplemented
}

func (r *repositoryImpl) DeleteMessageReaction(messageID uuid.UUID) (*domain.MessageReaction, error) {
	return nil, domain.ErrNotImplemented
}
