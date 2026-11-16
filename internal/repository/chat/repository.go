package chat

import (
	"context"
	"fmt"

	repo "github.com/pavel-pj/go-chat/internal/repository"
	desc "github.com/pavel-pj/go-chat/pkg/chat_v1"
	db "github.com/pavel-pj/platform_lib/pkg/client/db"

	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "chats"

	idColumn        = "id"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type ChatRepository struct {
	db db.Client
}

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func NewChatRepository(db db.Client) repo.ChatRepository {
	return &ChatRepository{
		db: db,
	}
}

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

/*
func (r *Repository) Delete(ctx context.Context, id int64) error {

	// first delete rows from chat_users
	sqlStr, args, err := psql.
		Delete("chat_users").
		Where(sq.Eq{"chat_id": id}).ToSql()
	if err != nil {
		return fmt.Errorf("failed to build delete-chat_users query: %w", err)
	}

	tags, err := r.DB.Exec(ctx, sqlStr, args...)
	if err != nil {
		return fmt.Errorf("failed to delete chat_users : %w", err)
	}
	if tags.RowsAffected() == 0 {
		return fmt.Errorf("zero rows of chat_users deleted : %w", err)
	}

	sqlStr, args, err = psql.
		Delete("chats").
		Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return fmt.Errorf("failed to build delete-chat query: %w", err)
	}

	tags, err = r.DB.Exec(ctx, sqlStr, args...)
	if err != nil {
		return fmt.Errorf("failed to delete chat : %w", err)
	}
	if tags.RowsAffected() == 0 {
		return fmt.Errorf("zero rows deleted : %w", err)
	}

	return nil

}

func (r *Repository) Get(ctx context.Context, id int64) (*desc.Chat, error) {

	sqlStr, args, err := psql.Select().
		From(tableName).
		Columns(idColumn, createdAtColumn, updatedAtColumn).
		Where(sq.Eq{idColumn: id}).
		ToSql()

	if err != nil {

		return nil, fmt.Errorf("failed to build get-chat query: %w", err)

	}
	var chat model.Chat
	err = r.DB.QueryRow(ctx, sqlStr, args...).
		Scan(
			&chat.ID,
			&chat.CreatedAt,
			&chat.UpdatedAt,
		)
	if err != nil {
		return nil, fmt.Errorf("could not getting data from chats table query: %w", err)

	}

	return converter.ToChatFromRepo(&chat), nil
}
*/
