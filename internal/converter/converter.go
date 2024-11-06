package converter

import (
	"chat-server/internal/entities"
	desc "github.com/Chigvero/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToSendMessage(m desc.SendMessageRequest) entities.SendMessage {
	return entities.SendMessage{
		FromUserName: m.From,
		MessageText:  m.Text,
		Timestamp:    m.GetTimestamp().AsTime(),
	}
}

func ToDescSendMessage(m entities.SendMessage) desc.SendMessageRequest {
	return desc.SendMessageRequest{
		From:      m.FromUserName,
		Text:      m.MessageText,
		Timestamp: timestamppb.New(m.Timestamp),
	}
}
