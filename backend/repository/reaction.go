package repository

import (
	"fmt"

	"time"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_09/domain"
)

type MessageReactionRepository interface {
	DeleteMessageReaction(messageID uuid.UUID, username string) error
	GetMessageReaction(messageID uuid.UUID, username string) (*domain.MessageReaction, error)
	GetReactionsToMessage(messageID uuid.UUID) ([]*domain.MessageReaction, error)
	InsertMessageReaction(messageID uuid.UUID, username string) (*domain.MessageReaction, error)
}

func (r *repositoryImpl) GetMessageReaction(messageID uuid.UUID, username string) (*domain.MessageReaction, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*)  FROM message_reactions WHERE Message_id = ?", messageID)
	if err != nil {
		return &domain.GetMessageReactionResponse{}, err
	}

	var exists bool
	err = r.db.Get(&exists, "SELECT EXISTS (SELECT 1 FROM message_reactions WHERE message_id = ?)", messageID)
	if err != nil {
		return &domain.GetMessageReactionResponse{}, err
	}

	res := &domain.GetMessageReactionResponse{
		Count:              count,
		UserAlreadyReacted: exists,
	}
	return res, nil
}

func (r *repositoryImpl) InsertMessageReaction(messageID uuid.UUID, username string) (*domain.MessageReaction, error) {
	now := time.Now()
	_, err := r.db.Exec("INSERT INTO message_reactions (message_id,username,created_at)VALUES(?,?,?)", messageID, username, now)
	if err != nil {
		return &domain.InsertMessageReactionResponce{}, err
	}

	var count int
	err = r.db.Get(&count, "SELECT COUNT(*) FROM message_reactions WHERE Message_id = ?", messageID)
	if err != nil {
		return &domain.InsertMessageReactionResponce{}, err
	}
	res := &domain.InsertMessageReactionResponce{
		MessageID: messageID,
		UserName:  username,
		CreatedAt: now,
		GetMessageReactionResponse: domain.GetMessageReactionResponse{
			Count: count,
		}}
	return res, nil
}

func (r *repositoryImpl) DeleteMessageReaction(messageID uuid.UUID, username string) error {
	res, err := r.db.Exec("DELETE FROM message_reactions WHERE message_id = ?", messageID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println("Rows Affected:", rowsAffected)
	return err
}

func (r *repositoryImpl) GetReactionsToMessage(messageID uuid.UUID) ([]domain.MessageReaction, error) {
	return nil, domain.ErrNotImplemented
}
