package errl

import (
	"context"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
)

func ListenForErrors(ctx context.Context, errs <-chan error) error {
	logrus.Infof("starting error listener")

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context done")
		case err := <-errs:
			log.Printf("%v", err)
		}
	}
}
