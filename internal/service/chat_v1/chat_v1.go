package chat_v1

import (
	"chat-server/internal/repository"
	"context"
)

type ChatService struct {
	repo repository.ChatRepository
}

func NewChatSevice(repo repository.ChatRepository) *ChatService {
	return &ChatService{
		repo: repo,
	}
}

func (s *ChatService) Create(ctx context.Context, usrs []string) (int64, error) {
	return s.repo.Create(ctx, usrs)
}
func (s *ChatService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
func (s *ChatService) SendMessage(ctx context.Context) {}
