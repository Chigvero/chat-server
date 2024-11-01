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

func (s *ChatService) Create(ctx context.Context, usrs []string) {}
func (s *ChatService) Delete(ctx context.Context)                {}
func (s *ChatService) SendMessage(ctx context.Context)           {}
