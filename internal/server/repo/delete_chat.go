package repo

import "context"

func (r *Repository) DeleteChat(ctx context.Context, chatID int64) error {
	query := `
		delete from chats
		where chat_id = $1`

	_, err := r.pool.Exec(ctx, query, chatID)
	return err
}
