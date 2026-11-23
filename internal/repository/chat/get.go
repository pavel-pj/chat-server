package chat

import (
	"context"
	"fmt"

	"github.com/pavel-pj/go-chat/internal/model"
	"github.com/pavel-pj/go-chat/internal/repository/chat/converter"
	repoModel "github.com/pavel-pj/go-chat/internal/repository/chat/model"
	"github.com/pavel-pj/platform_lib/pkg/client/db"

	sq "github.com/Masterminds/squirrel"
)

func (r *ChatRepository) Get(ctx context.Context, id int64) (model.Chat, error) {

	sqlStr, args, err := psql.Select().
		From(tableName).
		Columns(idColumn, createdAtColumn, updatedAtColumn).
		Where(sq.Eq{idColumn: id}).
		ToSql()

	if err != nil {

		return model.Chat{}, fmt.Errorf("failed to build get-chat query: %w", err)

	}

	q := db.Query{
		Name:     "chat_repository.get",
		QueryRow: sqlStr,
	}

	var chat repoModel.Chat
	err = r.db.DB().ScanOneContext(ctx, &chat, q, args...)

	if err != nil {
		return model.Chat{}, fmt.Errorf("could not getting data from chats table query: %w", err)

	}

	return converter.FromChatRepoToService(chat), nil
}
