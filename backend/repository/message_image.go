package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/traP-jp/h25s_09/domain"
)

type MessageImageRepository interface {
	GetMessageImage(imageID uuid.UUID) (*domain.MessageImage, error)
	CreateMessageImage(messageID uuid.UUID, data []byte, mime string) (*domain.MessageImage, error)
}

type messageImageRepositoryImpl struct {
	db *sqlx.DB
}

func NewMessageImageRepository(db *sqlx.DB) MessageImageRepository {
	return &messageImageRepositoryImpl{db: db}
}

func (r *messageImageRepositoryImpl) GetMessageImage(imageID uuid.UUID) (*domain.MessageImage, error) {
	return nil, domain.ErrNotImplemented
}

func (r *messageImageRepositoryImpl) CreateMessageImage(messageID uuid.UUID, data []byte, mime string) (*domain.MessageImage, error) {
	return nil, domain.ErrNotImplemented
}
