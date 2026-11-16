package converter

import (
	//"chat-server/internal/grpc_server/repository/chat/model"
	model "github.com/pavel-pj/go-chat/internal/repository/chat/model"
	desc "github.com/pavel-pj/go-chat/pkg/chat_v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToChatFromRepo(chat *model.Chat) *desc.Chat {
	var updatedAd *timestamppb.Timestamp

	if chat.UpdatedAt.Valid {
		updatedAd = timestamppb.New(chat.UpdatedAt.Time)
	}

	var result desc.Chat
	result = desc.Chat{
		Id:        chat.ID,
		CreatedAt: timestamppb.New(chat.CreatedAt),
		UpdatedAt: updatedAd,
	}

	return &result

}
