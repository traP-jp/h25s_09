package repository

import "github.com/jmoiron/sqlx"

type Repository interface {
	AchievementsRepository
	MessageRepository
	MessageReactionRepository
	MessageImageRepository
}

type repositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{db: db}
}
