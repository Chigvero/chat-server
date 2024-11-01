package repository

import (
	"chat-server/internal/repository/postgres/chat_v1"
	"context"
	"github.com/jackc/pgx/v5"
)

type ChatRepository interface {
	Create(ctx context.Context)
	Delete(ctx context.Context)
	SendMessage(ctx context.Context)
}

type Repository struct {
	ChatRepository
}

func NewRepository(DB *pgx.Conn) *Repository {
	return &Repository{
		ChatRepository: chat_v1.NewChatPostgres(DB),
	}
}
