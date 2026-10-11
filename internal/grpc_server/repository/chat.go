package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func NewRepository(DB *pgxpool.Pool) Repository {
	return Repository{
		DB: DB,
	}
}

func (r *Repository) Store(ctx context.Context, users []string) (int64, error) {

	var chatID int64
	err := r.DB.QueryRow(ctx, `INSERT INTO chats DEFAULT VALUES RETURNING id`).Scan(&chatID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert chat: %w", err)
	}

	insert := psql.Insert("chat_users").Columns("chat_id", "user_name")

	fmt.Println(users)
	for _, user := range users {
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

	return chatID, nil
}

func (r *Repository) Delete(ctx context.Context, id int64) error {

	// first delete rows from chat_users
	sqlSrt, args, err := psql.
		Delete("chat_users").
		Where(sq.Eq{"chat_id": id}).ToSql()
	if err != nil {
		return fmt.Errorf("failed to build delete-chat_users query: %w", err)
	}

	tags, err := r.DB.Exec(ctx, sqlSrt, args...)
	if err != nil {
		return fmt.Errorf("failed to delete chat_users : %w", err)
	}
	if tags.RowsAffected() == 0 {
		return fmt.Errorf("zero rows of chat_users deleted : %w", err)
	}

	sqlSrt, args, err = psql.
		Delete("chats").
		Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return fmt.Errorf("failed to build delete-chat query: %w", err)
	}

	tags, err = r.DB.Exec(ctx, sqlSrt, args...)
	if err != nil {
		return fmt.Errorf("failed to delete chat : %w", err)
	}
	if tags.RowsAffected() == 0 {
		return fmt.Errorf("zero rows deleted : %w", err)
	}

	return nil

}
