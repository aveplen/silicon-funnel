package repo

import (
	"context"
)

func (r *Repository) InsertChat(ctx context.Context, tgChatID int64) error {
	query := `
		insert into chats (
			tg_chat_id
		) values (
			$1
		)`

	if _, err := r.pool.Exec(ctx, query, tgChatID); err != nil {
		return err
	}

	return nil
}
