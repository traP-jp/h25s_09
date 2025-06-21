package repository

import (
	"github.com/google/uuid"

	"github.com/traP-jp/h25s_09/domain"
)

type MessageRepository interface {
	CreateMessage(author, content string, parentID uuid.UUID) (*domain.Message, error)
	GetMessageByID(id uuid.UUID) (*domain.Message, error)
}

func (r *repositoryImpl) CreateMessage(author, content string, parentID uuid.UUID) (*domain.Message, error) {
	
	message := &domain.Message{
        ID:       uuid.New(),
        Author:   author,
        Content:  content,
        ParentID: parentID,
    }

	// データベースに保存
	_, err := r.db.Exec("INSERT INTO messages (id, author, message, replies_id) VALUES (?, ?, ?, ?)",
		message.ID, message.Author, message.Content, message.ParentID,
	)
	
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (r *repositoryImpl) GetMessageByID(id uuid.UUID) (*domain.Message, error) {


	message := &domain.Message{
		ID:       id,
		Author:   "",
		Content:  "",
		ParentID: uuid.Nil,
	}

	// データベースからメッセージを取得しmessageに格納
	err := r.db.Get(message, "SELECT id, author, message, replies_id FROM messages WHERE id = ?", id)
	if err != nil {
		return nil, err // エラーが発生した場合はnilを返す
	}

	return message, nil
}