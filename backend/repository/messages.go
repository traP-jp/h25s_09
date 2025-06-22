package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/traP-jp/h25s_09/domain"
)

type MessageRepository interface {
	CreateMessage(author, content string, parentID uuid.UUID) (*domain.Message, error)
	GetMessageByID(id uuid.UUID) (*domain.Message, error)
	GetMessages(limit, offset int64, username string, includeReplies bool) ([]domain.Message, error)
	GetRepliesByMessageID(messageID uuid.UUID) ([]*domain.Message, error)
}

type Message struct {
	ID        uuid.UUID `db:"id"`
	Author    string    `db:"author"`
	Content   string    `db:"message"`
	ParentID  uuid.UUID `db:"replies_to"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (r *repositoryImpl) GetMessages(limit, offset int64, username string, includeReplies bool) ([]domain.Message, error) {
	var messages []Message
	query := "SELECT id, author, message, replies_to, created_at, updated_at FROM messages"
	args := []any{}

	if username != "" && includeReplies {
		query += " WHERE author = ?"
		args = append(args, username)
	}
	if username != "" && !includeReplies {
		query += " WHERE author = ? AND replies_to = ?"
		args = append(args, username, uuid.Nil)
	}
	if username == "" && includeReplies {
		query += " WHERE replies_to = ?"
		args = append(args, uuid.Nil)
	}

	query += " ORDER BY created_at DESC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	err := r.db.Select(&messages, query, args...)
	if err != nil {
		return nil, err
	}

	// domain.Messageに変換して返す
	domainMessages := make([]domain.Message, len(messages))
	for i, msg := range messages {
		domainMessages[i] = domain.Message{
			ID:        msg.ID,
			Author:    msg.Author,
			Content:   msg.Content,
			ParentID:  msg.ParentID,
			CreatedAt: msg.CreatedAt,
			UpdatedAt: msg.UpdatedAt,
		}
	}

	return domainMessages, nil
}

func (r *repositoryImpl) CreateMessage(author, content string, parentID uuid.UUID) (*domain.Message, error) {
	message := &domain.Message{
		ID:       uuid.Must(uuid.NewV7()),
		Author:   author,
		Content:  content,
		ParentID: parentID,
	}

	// データベースに保存
	_, err := r.db.Exec("INSERT INTO messages (id, author, message, replies_to) VALUES (?, ?, ?, ?)",
		message.ID, message.Author, message.Content, message.ParentID,
	)
	if err != nil {
		return nil, err
	}

	return r.GetMessageByID(message.ID)
}

func (r *repositoryImpl) GetMessageByID(id uuid.UUID) (*domain.Message, error) {
	var message Message

	// データベースからメッセージを取得しmessageに格納
	err := r.db.Get(&message, "SELECT id, author, message, replies_to, created_at, updated_at FROM messages WHERE id = ?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err // エラーが発生した場合はnilを返す
	}

	// domain.Messageに変換して返す
	return &domain.Message{
		ID:        message.ID,
		Author:    message.Author,
		Content:   message.Content,
		ParentID:  message.ParentID,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
	}, nil
}

func (r *repositoryImpl) GetRepliesByMessageID(messageID uuid.UUID) ([]*domain.Message, error) {
	var replies []*Message
	err := r.db.Select(&replies, "SELECT * FROM messages WHERE replies_to = ? ORDER BY created_at DESC", messageID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err // エラーが発生した場合はnilを返す
	}

	// Messageをdomain.Messageに変換
	domainReplies := make([]*domain.Message, len(replies))
	for i := range replies {
		domainReplies[i] = &domain.Message{
			ID:        replies[i].ID,
			Author:    replies[i].Author,
			Content:   replies[i].Content,
			ParentID:  replies[i].ParentID,
			CreatedAt: replies[i].CreatedAt,
			UpdatedAt: replies[i].UpdatedAt,
		}
	}

	return domainReplies, nil
}
