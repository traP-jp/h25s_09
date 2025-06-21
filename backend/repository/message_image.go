package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_09/domain"
)

type MessageImageRepository interface {
	GetMessageImage(imageID uuid.UUID) (*domain.MessageImage, error)
	CreateMessageImage(messageID uuid.UUID, data []byte, mime string) (*domain.MessageImage, error)
	GetMessageImageIDByMessageID(messageID uuid.UUID) (uuid.UUID, error)
}

type repoMessageImage struct {
	Id        uuid.UUID `db:"id"`
	MessageID uuid.UUID `db:"message_id"`
	Data      []byte    `db:"data"`
	Mime      string    `db:"mime"`
	CreatedAt time.Time `db:"created_at"`
}

func (r *repositoryImpl) GetMessageImage(imageID uuid.UUID) (*domain.MessageImage, error) {
	var img repoMessageImage
	err := r.db.Get(&img, "SELECT * FROM message_images WHERE id=?", imageID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		} else {
			return nil, err
		}
	}
	return &domain.MessageImage{
		ID:        img.Id,
		MessageID: img.MessageID,
		Data:      img.Data,
		Mime:      img.Mime,
		CreatedAt: img.CreatedAt,
	}, nil
}

func (r *repositoryImpl) CreateMessageImage(messageID uuid.UUID, data []byte, mime string) (*domain.MessageImage, error) {
	img := repoMessageImage{
		Id:        uuid.Must(uuid.NewV7()),
		MessageID: messageID,
		Data:      data,
		Mime:      mime,
	}
	res, err := r.db.NamedExec("INSERT INTO message_images (id, message_id, data, mime) VALUES (:id, :message_id, :data, :mime)", img)
	if err != nil {
		return nil, err
	}
	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return nil, errors.New("no rows affected")
	}
	return r.GetMessageImage(img.Id)
}

func (r *repositoryImpl) GetMessageImageIDByMessageID(messageID uuid.UUID) (uuid.UUID, error) {
	var imageId uuid.UUID
	err := r.db.Get(&imageId, "SELECT id FROM message_images WHERE message_id=?", messageID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.Nil, domain.ErrNotFound
		}
		return uuid.Nil, err
	}
	return imageId, nil
}
