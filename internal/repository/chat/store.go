package chat

import (
	"context"
	"fmt"

	desc "github.com/pavel-pj/go-chat/pkg/chat_v1"
	db "github.com/pavel-pj/platform_lib/pkg/client/db"

	sq "github.com/Masterminds/squirrel"
)

func (r *ChatRepository) Store(ctx context.Context, users *desc.ChatInfo) (int64, error) {

	sqlStr, args, err := psql.Insert(tableName).
		Columns(idColumn).
		Values(sq.Expr("DEFAULT")).
		Suffix("RETURNING ID").
		ToSql()

	q := db.Query{
		Name:     "chat_repository.store",
		QueryRow: sqlStr,
	}

	var id int64
	err = r.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, fmt.Errorf("failed to insert chat: %w", err)
	}
	/*

		insert := psql.Insert("chat_users").Columns("chat_id", "user_name")

		fmt.Println(users)
		for _, user := range users.Usernames {
			insert = insert.Values(chatID, user)
		}

		sqlStr, args, err := insert.ToSql()
		if err != nil {
			return 0, fmt.Errorf("failed to insert users: %w", err)
		}

		_, err = r.DB.Exec(ctx, sqlStr, args...)
		if err != nil {
			return 0, fmt.Errorf("failed to build chat_users query: %w", err)
		}
	*/
	return id, nil
}
