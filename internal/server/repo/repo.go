package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/aveplen/silicon-funnel/internal/server/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrNotFound = errors.New("no rows found")

type Repository struct {
	pool  *pgxpool.Pool
	bumps chan *OffsetBump
	errs  chan error
}

func NewRepository(ctx context.Context, cfg *config.Config, errs chan error) (*Repository, error) {

	connStringFormat := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	connString := fmt.Sprintf(
		connStringFormat,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
	)

	pool, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &Repository{
		pool:  pool,
		bumps: make(chan *OffsetBump, 100),
		errs:  errs,
	}, nil
}
