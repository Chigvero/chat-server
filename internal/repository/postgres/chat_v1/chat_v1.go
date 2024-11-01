package chat_v1

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type ChatPostgres struct {
	DB *pgx.Conn
}

func NewChatPostgres(DB *pgx.Conn) *ChatPostgres {
	return &ChatPostgres{
		DB: DB,
	}
}

func (r *ChatPostgres) Create(ctx context.Context) {

}

func (r *ChatPostgres) Delete(ctx context.Context) {

}
func (r *ChatPostgres) SendMessage(ctx context.Context) {

}
