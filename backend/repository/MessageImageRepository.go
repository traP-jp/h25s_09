package repository

import (
	"github.com/google/uuid"
	"github.com/traP-jp/h25s_09/domain"
)

type MessageImageRepository interface {
	GetMessageImage(imageID uuid.UUID) (domain.MessageImage, error)
	CreateMessageImage(request domain.PostMessageImageRequest) (domain.MessageImage, error)
}
