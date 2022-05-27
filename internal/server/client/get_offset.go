package client

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"github.com/emersion/go-imap/client"
	"github.com/sirupsen/logrus"
)

const errFormat = "could not fetch offset: %w"

func GetOffset(ctx context.Context, mailbox *pb.MailboxV1) (uint32, error) {
	logrus.Infof("fetching initial offset for %v", mailbox)

	offsetChan := make(chan uint32)
	errChan := make(chan error)

	go func() {
		address := fmt.Sprintf("%s:%d", mailbox.Host, mailbox.Port)

		c, err := client.DialTLS(address, nil)
		if err != nil {
			errChan <- fmt.Errorf(errFormat, err)
		}

		if err := c.Login(mailbox.Username, mailbox.Password); err != nil {
			errChan <- fmt.Errorf(errFormat, err)
		}

		fetched, err := c.Select(mailbox.Mailbox, true)
		if err != nil {
			errChan <- fmt.Errorf(errFormat, err)
		}

		offsetChan <- fetched.Messages
	}()

	select {
	case <-ctx.Done():
		return 0, fmt.Errorf(errFormat, errors.New("context done"))

	case offset := <-offsetChan:
		return offset, nil

	case err := <-errChan:
		return 0, err
	}
}
