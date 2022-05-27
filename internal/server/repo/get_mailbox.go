package repo

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

func (r *Repository) GetMailboxByID(ctx context.Context, mailboxID int64) (*pb.MailboxV1, error) {
	query := `
		select 
			mailbox_id,
			name,
			host,
			port,
			username,
			password
		from mailboxes
		where mailbox_id = $1`

	var mailbox pb.MailboxV1

	row := r.pool.QueryRow(ctx, query, mailboxID)
	if err := row.Scan(
		&mailbox.MailboxID,
		&mailbox.Mailbox,
		&mailbox.Host,
		&mailbox.Port,
		&mailbox.Username,
		&mailbox.Password,
	); err != nil {
		return nil, err
	}

	return &mailbox, nil
}
