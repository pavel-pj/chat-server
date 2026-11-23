package chat

import (
	repo "github.com/pavel-pj/go-chat/internal/repository"
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
