package repository

import (
	"github.com/traP-jp/h25s_09/domain"
)

type MessageImageRepository interface {
	GetMessageImage() (domain.MessageImage, error)
	CreateMessageImage(req domain.PostMessageImageRequest) (domain.MessageImage, error)
}
