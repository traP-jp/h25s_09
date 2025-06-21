package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/traP-jp/h25s_09/domain"
)

type MessageRepository interface {
	CreateMessage(author, content string, parentID uuid.UUID) (*domain.Message, error)
	GetMessageByID(id uuid.UUID) (*domain.Message, error)
}

type messageRepositoryImpl struct {
	db *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) MessageRepository {
	return &messageRepositoryImpl{db: db}
}

func (r *messageRepositoryImpl) CreateMessage(author, content string, parentID uuid.UUID) (*domain.Message, error) {
	return nil, domain.ErrNotImplemented
}

func (r *messageRepositoryImpl) GetMessageByID(id uuid.UUID) (*domain.Message, error) {
	return nil, domain.ErrNotImplemented
}
