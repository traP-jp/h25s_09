package repository

import (
	"errors"
	"database/sql"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_09/domain"
)

type RepliesRepository interface {
	GetRepliesByMessageID(messageID uuid.UUID) ([]*domain.Message, error)
}

func (r *repositoryImpl) GetRepliesByMessageID(messageID uuid.UUID) ([]*domain.Message, error) {
	var replies []domain.Message
	err := r.db.Select(&replies, "SELECT id, author, message, replies_id, created_at, updated_at FROM messages WHERE replies_id = ?", messageID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err // エラーが発生した場合はnilを返す
	}

	// 今はスライスなので、ポインタのスライスに変換
	replyPointers := make([]*domain.Message, len(replies))
	for i := range replies {
		replyPointers[i] = &replies[i]
	}
	return replyPointers, nil
}