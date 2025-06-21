package repository

import (
	"time"
	"github.com/google/uuid"

	"github.com/traP-jp/h25s_09/domain"
)

type MessageRepository interface {
	CreateMessage(author, content string, parentID uuid.UUID) (*domain.Message, error)
	GetMessageByID(id uuid.UUID) (*domain.Message, error)
}

type Message struct {
    ID        uuid.UUID `db:"id"`
    Author    string    `db:"author"`
    Content   string    `db:"content"`
    ParentID  uuid.UUID `db:"parent_id"`
    CreatedAt time.Time `db:"created_at"`
    UpdatedAt time.Time `db:"updated_at"`
}

func (r *repositoryImpl) CreateMessage(author, content string, parentID uuid.UUID) (*domain.Message, error) {
	message := &domain.Message{
		ID:        uuid.Must(uuid.NewV7()),
		Author:    author,
		Content:   content,
		ParentID:  parentID,
	}

	// データベースに保存
	_, err := r.db.Exec("INSERT INTO messages (id, author, message, replies_id) VALUES (?, ?, ?, ?)",
		message.ID, message.Author, message.Content, message.ParentID,
	)
	if err != nil {
		return nil, err
	}

	return r.GetMessageByID(message.ID)
}

func (r *repositoryImpl) GetMessageByID(id uuid.UUID) (*domain.Message, error) {
	message := &Message{
		ID:       id,
		Author:   "",
		Content:  "",
		ParentID: uuid.Nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	// データベースからメッセージを取得しmessageに格納
	err := r.db.Get(message, "SELECT id, author, message, replies_id, created_at, updated_at FROM messages WHERE id = ?", id)
	if err != nil {
		return nil, err // エラーが発生した場合はnilを返す
	}

	// domain.Messageに変換して返す
	return &domain.Message{
		ID:       message.ID,
		Author:   message.Author,
		Content:  message.Content,
		ParentID: message.ParentID,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
	}, nil
}