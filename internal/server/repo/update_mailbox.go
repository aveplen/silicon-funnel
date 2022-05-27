package repo

import (
	"context"

	"github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

func (r *Repository) UpdateMailbox(ctx context.Context, mailbox *imap_concentrator.MailboxV1) error {
	query := `
		update mailboxes set (
			name = $1,
			host = $2,
			port = $3,
			username = $4,
			password = $5
		)
		mailbox_id = $6`

	if _, err := r.pool.Exec(ctx, query,
		mailbox.Mailbox,
		mailbox.Host,
		mailbox.Port,
		mailbox.Username,
		mailbox.Password,
		mailbox.MailboxID,
	); err != nil {
		return err
	}

	return nil
}
