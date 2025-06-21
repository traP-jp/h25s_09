package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID         uuid.UUID  `json:"id"`
	Author     string     `json:"author"`
	Content    string     `json:"content"`
	ImageID    *uuid.UUID `json:"imageId"`
	Reactions  Reactions  `json:"reactions"`
	ReplyCount int        `json:"replyCount"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
}

type Reactions struct {
	Count      int  `json:"count"`
	MyReaction bool `json:"myReaction"`
}
