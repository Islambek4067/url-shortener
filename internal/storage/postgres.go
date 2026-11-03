package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	pool *pgxpool.Pool
}

func NewPostgres(ctx context.Context, databaseURL string) (*Postgres, error) {
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, err
	}

	return &Postgres{pool: pool}, nil
}

func (p *Postgres) Save(code, curl string) error {
	_, err := p.pool.Exec(context.Background(), "insert into links (code, original) values ($1, $2)", code, curl)
	return err
}

func (p *Postgres) Get(code string) (Link, bool, error) {
	var l Link
	err := p.pool.QueryRow(context.Background(), "SELECT original, clicks, created_at FROM links WHERE code = $1", code).Scan(&l.Original, &l.Clicks, &l.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return Link{}, false, nil
	}
	if err != nil {
		return Link{}, false, nil
	}
	return l, true, nil
}

func (p *Postgres) AddClick(code string) error {
	_, err := p.pool.Exec(context.Background(), "UPDATE links SET clicks = clicks +1 WHERE code =  $1", code)
	if err != nil {
		return err
	}
	return nil
}
