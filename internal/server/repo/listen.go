package repo

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

func (r *Repository) ListenForBumps(ctx context.Context) error {
	logrus.Infof("starting offset bump listener")

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context done")

		case bump := <-r.bumps:
			if err := r.BumpOffset(ctx, bump.MailboxID, bump.Bump); err != nil {
				r.errs <- err
			}
		}
	}
}
