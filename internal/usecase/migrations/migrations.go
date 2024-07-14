package migrations

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
)

func Up(pool *pgxpool.Pool, migrationsPath string) error {
	db := stdlib.OpenDBFromPool(pool)

	if err := goose.SetDialect("postgres"); err != nil {
		return errors.Wrap(err, "failed to set dialect")
	}

	if err := goose.Up(db, migrationsPath); err != nil {
		return errors.Wrap(err, "failed to run migrations")
	}

	return nil
}
