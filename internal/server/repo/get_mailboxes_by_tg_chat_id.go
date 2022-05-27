package repo

import (
	"context"
	"errors"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"github.com/jackc/pgx/v4"
)

func (r *Repository) GetMailboxesByTgChatID(ctx context.Context, tgChatID int64) ([]*pb.MailboxV1, []MailboxMeta, error) {
	query := `
		select 
			chats.tg_chat_id,
			mailboxes.mailbox_id,
			mailboxes.name,
			mailboxes.host,
			mailboxes.port,
			mailboxes.username,
			mailboxes.password,
			mailboxes.offset
		from chats
			join mailboxes on chats.chat_id = mailboxes.chat_id
		where 
			chats.tg_chat_id = $1`

	rows, err := r.pool.Query(ctx, query, tgChatID)
	if errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
		return nil, nil, err
	}

	var mailboxes []*pb.MailboxV1
	var metas []MailboxMeta

	for rows.Next() {
		mailbox := &pb.MailboxV1{}
		var meta MailboxMeta

		if err := rows.Scan(
			&meta.TgChatID,
			&mailbox.MailboxID,
			&mailbox.Mailbox,
			&mailbox.Host,
			&mailbox.Port,
			&mailbox.Username,
			&mailbox.Password,
			&meta.Offset,
		); err != nil {
			return nil, nil, err
		}

		mailboxes = append(mailboxes, mailbox)
		metas = append(metas, meta)
	}

	return mailboxes, metas, nil
}
