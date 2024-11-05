package chat_v1

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

const (
	chatsTable     = "chats"
	chatUsersTable = "chatUsers"
)

type ChatPostgres struct {
	DB *pgx.Conn
}

func NewChatPostgres(DB *pgx.Conn) *ChatPostgres {
	return &ChatPostgres{
		DB: DB,
	}
}

func (r *ChatPostgres) Create(ctx context.Context, usrs []string) (int64, error) {
	tx, err2 := r.DB.Begin(ctx)
	if err2 != nil {
		return 0, err2
	}
	var id int64
	createChatQuery := fmt.Sprintf("INSERT INTO %s (created_at) VALUES (DEFAULT) RETURNING id;\n", chatsTable)
	err := tx.QueryRow(ctx, createChatQuery).Scan(&id)
	if err != nil {
		err2 = tx.Rollback(ctx)
		if err2 != nil {
			return 0, err2
		}
		return 0, err
	}

	insertChatUserExec := fmt.Sprintf("INSERT INTO %s(chat_id,user_name) VALUES($1,$2)", chatUsersTable)
	for _, usr := range usrs {
		rowsAffected, err := tx.Exec(ctx, insertChatUserExec, id, usr)
		if err != nil {
			err2 = tx.Rollback(ctx)
			if err2 != nil {
				return 0, err2
			}
			return 0, err
		}
		log.Println(":", rowsAffected)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ChatPostgres) Delete(ctx context.Context, id int64) error {
	tx, err2 := r.DB.Begin(ctx)
	if err2 != nil {
		return err2
	}
	deleteChatExec := fmt.Sprintf("DELETE FROM %s WHERE id=$1", chatsTable)
	_, err := tx.Exec(ctx, deleteChatExec, id)
	if err != nil {
		err2 = tx.Rollback(ctx)
		if err2 != nil {
			return err2
		}
		return err
	}
	deleteChatUsersExec := fmt.Sprintf("DELETE FROM %s id=$1", chatUsersTable)
	_, err = tx.Exec(ctx, deleteChatUsersExec, id)
	if err != nil {
		err2 = tx.Rollback(ctx)
		if err2 != nil {
			return err2
		}
		return err
	}
	err2 = tx.Commit(ctx)
	if err2 != nil {
		return err2
	}
	return nil
}
func (r *ChatPostgres) SendMessage(ctx context.Context) {

}
