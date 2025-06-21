package repository

import (
	"github.com/google/uuid"
	"github.com/traP-jp/h25s_09/domain"
)

type MessageImageRepository interface {
	GetMessageImage(imageID uuid.UUID) (domain.MessageImage, error)
	CreateMessageImage(messageID uuid.UUID, data []byte, mime string) (domain.MessageImage, error)
}
