package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresClient struct {
	*pgxpool.Pool
	cfg *pgxpool.Config
}

func NewClient(connString string) (*PostgresClient, error) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config postgres: %w", err)
	}

	cfg.MaxConns = 5
	cfg.MinConns = 3
	cfg.MaxConnIdleTime = 15 * time.Second
	cfg.HealthCheckPeriod = 600 * time.Millisecond

	return &PostgresClient{
		cfg: cfg,
	}, nil
}

func (c *PostgresClient) Connect(ctx context.Context) error {
	pool, err := pgxpool.NewWithConfig(ctx, c.cfg)
	if err != nil {
		return fmt.Errorf("unable to create new connection to %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("unable to connect to ping postgres: %w", err)
	}
	c.Pool = pool

	return nil
}

func (c *PostgresClient) Close(ctx context.Context) error {
	c.Pool.Close()

	return nil
}
