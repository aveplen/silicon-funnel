package repo

import (
	"context"
)

type OffsetBump struct {
	MailboxID int64
	Bump      int
}

func (r *Repository) PostBump(bump *OffsetBump) {
	r.bumps <- bump
}

func (r *Repository) BumpOffset(ctx context.Context, mailboxID int64, bump int) error {
	query := `
		update mailboxes set 
		"offset" = "offset" + $1
		where mailbox_id = $2`

	if _, err := r.pool.Exec(ctx, query, bump, mailboxID); err != nil {
		return err
	}

	return nil
}
