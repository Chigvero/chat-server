package repository

import (
	"chat-server/internal/entities"
	"chat-server/internal/repository/postgres/chat_v1"
	"context"
	"github.com/jackc/pgx/v5"
)

type ChatRepository interface {
	Create(ctx context.Context, usrs []string) (int64, error)
	Delete(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, message entities.SendMessage) error
}

type Repository struct {
	ChatRepository
}

func NewRepository(DB *pgx.Conn) *Repository {
	return &Repository{
		ChatRepository: chat_v1.NewChatPostgres(DB),
	}
}
