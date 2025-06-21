package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/traP-jp/h25s_09/domain"
)

type MessageRactionRepository interface {
	GetMessageReaction(messageID uuid.UUID) (*domain.MessageReaction, error)
	InsertMessageReaction(messageID uuid.UUID) (*domain.MessageReaction, error)
	DeleteMessageReaction(messageID uuid.UUID) (*domain.MessageReaction, error)
}

type messageReactionRepositoryImpl struct {
	db *sqlx.DB
}

func NewMessageReactionRepository(db *sqlx.DB) MessageRactionRepository {
	return &messageReactionRepositoryImpl{db: db}
}

func (r *messageReactionRepositoryImpl) GetMessageReaction(messageID uuid.UUID) (*domain.MessageReaction, error) {
	return nil, domain.ErrNotImplemented
}

func (r *messageReactionRepositoryImpl) InsertMessageReaction(messageID uuid.UUID) (*domain.MessageReaction, error) {
	return nil, domain.ErrNotImplemented
}

func (r *messageReactionRepositoryImpl) DeleteMessageReaction(messageID uuid.UUID) (*domain.MessageReaction, error) {
	return nil, domain.ErrNotImplemented
}
