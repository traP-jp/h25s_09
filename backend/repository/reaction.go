package repository

import (
	"github.com/google/uuid"
	"github.com/traP-jp/h25s_09/domain"
)

type MessageReactionRepository interface {
	DeleteMessageReaction(messageID uuid.UUID, username string) (*domain.MessageReaction, error)
	GetMessageReaction(messageID uuid.UUID, username string) (*domain.MessageReaction, error)
	GetReactionsToMessage(messageID uuid.UUID) ([]domain.MessageReaction, error)
	InsertMessageReaction(messageID uuid.UUID, username string) (*domain.MessageReaction, error)
}

func (r *repositoryImpl) DeleteMessageReaction(messageID uuid.UUID, username string) (*domain.MessageReaction, error) {
	return nil, domain.ErrNotImplemented
}

func (r *repositoryImpl) GetMessageReaction(messageID uuid.UUID, username string) (*domain.MessageReaction, error) {
	return nil, domain.ErrNotImplemented
}

func (r *repositoryImpl) GetReactionsToMessage(messageID uuid.UUID) ([]domain.MessageReaction, error) {
	return nil, domain.ErrNotImplemented
}

func (r *repositoryImpl) InsertMessageReaction(messageID uuid.UUID, username string) (*domain.MessageReaction, error) {
	return nil, domain.ErrNotImplemented
}
