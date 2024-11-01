package service

import (
	"chat-server/internal/repository"
	"chat-server/internal/service/chat_v1"
	"context"
)

type ChatService interface {
	Create(ctx context.Context, usrs []string)
	Delete(ctx context.Context)
	SendMessage(ctx context.Context)
}

type Service struct {
	ChatService
}

func NewService(repository repository.Repository) *Service {
	return &Service{
		ChatService: chat_v1.NewChatSevice(repository.ChatRepository),
	}
}
