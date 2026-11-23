package converter

import (
	"github.com/pavel-pj/go-chat/internal/model"
	desc "github.com/pavel-pj/go-chat/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromChatServiceToApi(chat model.Chat) *desc.GetResponse {

	return &desc.GetResponse{
		Chat: &desc.Chat{
			Id: chat.ID,
			Chat: &desc.ChatInfo{
				Usernames: []string{"test1", "test2"},
			},
			CreatedAt: timestamppb.New(chat.CreatedAt),
			UpdatedAt: timestamppb.New(chat.UpdatedAt),
		},
	}

}
