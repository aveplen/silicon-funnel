package repo

import (
	"context"
)

func (r *Repository) DeleteMailbox(ctx context.Context, mailboxID int64) error {
	query := `
		delete from mailboxes
		where
			mailboxID = $1`

	_, err := r.pool.Exec(ctx, query, mailboxID)
	return err
}
