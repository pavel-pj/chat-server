package chat

import (
	"github.com/pavel-pj/go-chat/internal/service"
	desc "github.com/pavel-pj/go-chat/pkg/chat_v1"
)

type ChatImplementation struct {
	desc.UnimplementedChatV1Server
	serv service.ChatService
}

func NewChatImplimentation(chatService service.ChatService) *ChatImplementation {
	return &ChatImplementation{
		serv: chatService,
	}
}
