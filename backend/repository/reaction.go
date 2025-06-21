package repository

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/traP-jp/h25s_09/domain"
)

type MessageReactionRepository interface {
	GetMessageReaction(messageID uuid.UUID) (int, error)
	InsertMessageReaction(messageID uuid.UUID,username string) (*domain.MessageReaction, error)
	DeleteMessageReaction(messageID uuid.UUID) (*domain.MessageReaction, error)
}

func (r *repositoryImpl) GetMessageReaction(messageID uuid.UUID) (int, error) {
	var count int
	err := r.db.Get(&count,"SELECT COUNT(*)  FROM message_reactions WHERE Message_id = ?",messageID)
	if err != nil {
		return 0 ,err
	}
	return count,nil
}

func (r *repositoryImpl) InsertMessageReaction(messageID uuid.UUID,username string) (*domain.MessageReaction, error) {
	res,err := r.db.Exec("INSERT INTO message_reactions (message_id,username,created_at)VLUES(?,?,0)",messageID,username)
	if err != nil || res == nil{
		return nil,err
	}
	return nil, domain.ErrNotImplemented
}

func (r *repositoryImpl) DeleteMessageReaction(messageID uuid.UUID) (*domain.MessageReaction, error) {
	res, err := r.db.Exec("DELETE FROM message_reactions WHERE message_id = ?",messageID)
	if err != nil {
    log.Fatal(err)  
}
rowsAffected, err := res.RowsAffected()  
if err != nil {
    log.Fatal(err) 
}
fmt.Println("Rows Affected:", rowsAffected)
return nil,nil
} 
