package converter

import (
	//"chat-server/internal/grpc_server/repository/chat/model"
	"github.com/pavel-pj/go-chat/internal/model"
	repoModel "github.com/pavel-pj/go-chat/internal/repository/chat/model"
)

func FromChatRepoToService(chat repoModel.Chat) model.Chat {

	return model.Chat{
		ID:        chat.ID,
		CreatedAt: chat.CreatedAt,
		UpdatedAt: chat.UpdatedAt,
	}

	/*
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
	*/
}
