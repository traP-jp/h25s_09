package repository

import (
	"database/sql"
	"errors"

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

type repoReaction struct {
	MessageID uuid.UUID `db:"message_id"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
}

func (r *repositoryImpl) GetMessageReaction(messageID uuid.UUID, username string) (*domain.MessageReaction, error) {
	var reaction repoReaction
	err := r.db.Get(&reaction, "SELECT * FROM message_reactions WHERE message_id=? AND username=?", messageID, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return &domain.MessageReaction{
		MessageID: reaction.MessageID,
		Username:  reaction.Username,
		CreatedAt: reaction.CreatedAt,
	}, nil
}

func (r *repositoryImpl) InsertMessageReaction(messageID uuid.UUID, username string) (*domain.MessageReaction, error) {
	res, err := r.db.Exec("INSERT INTO message_reactions (message_id, username) VALUES (?, ?)", messageID, username)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, domain.ErrConflict
	}
	return r.GetMessageReaction(messageID, username)
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
	if rowsAffected == 0 {
		return domain.ErrNotFound
	}
	return err
}

func (r *repositoryImpl) GetReactionsToMessage(messageID uuid.UUID) ([]*domain.MessageReaction, error) {
	var reactions []repoReaction
	err := r.db.Select(&reactions, "SELECT * FROM message_reactions WHERE message_id=?", messageID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*domain.MessageReaction{}, nil
		}
		return nil, err
	}
	result := make([]*domain.MessageReaction, len(reactions))
	for i, reaction := range reactions {
		result[i] = &domain.MessageReaction{
			MessageID: reaction.MessageID,
			Username:  reaction.Username,
			CreatedAt: reaction.CreatedAt,
		}
	}
	return result, nil
}
