package db

import (
	"context"
	"dev11/config"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	dbPool *pgxpool.Pool
}

// Database connection initialization
func DbInit(cfg config.Settings) *DB {
	db := DB{}
	// urlExample := "postgres://username:password@127.0.0.1:5432/database_name"
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", cfg.PgUser, cfg.PgPasswd, cfg.PgHost, cfg.PgBase)

	dbPool, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	db.dbPool = dbPool
	return &db
}
