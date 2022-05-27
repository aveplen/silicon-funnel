package repo

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

func (r *Repository) InsertMailbox(ctx context.Context, tgChatID int64, mailbox *pb.MailboxV1, offset uint32) error {
	query := `
		insert into mailboxes (
			chat_id,
			name,
			host,
			port,
			username,
			password,
			"offset"
		) values (
			(
				select chat_id
				from chats 
				where tg_chat_id = $1
			),
			$2, $3, $4, $5, $6, $7
		) returning mailbox_id`

	if err := r.pool.QueryRow(ctx, query,
		tgChatID,
		mailbox.Mailbox,
		mailbox.Host,
		mailbox.Port,
		mailbox.Username,
		mailbox.Password,
		offset,
	).Scan(
		&mailbox.MailboxID,
	); err != nil {
		return err
	}

	return nil
}
