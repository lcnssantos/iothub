package database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/lcnssantos/gomusic/config"
	_ "github.com/lib/pq"
)

func GetConnection() (*sql.DB, error) {
	configuration := config.Get()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", configuration.DB_HOST, configuration.DB_PORT, configuration.DB_USER, configuration.DB_PASS, configuration.DB_NAME)

	db, err := sql.Open(configuration.DB_DRIVER, connectionString)

	if err != nil {
		return nil, err
	}

	poolSize, err := strconv.Atoi(configuration.DB_POOL_SIZE)

	if err == nil {
		db.SetMaxOpenConns(poolSize)
	}

	return db, nil
}

func ExecuteTransaction(ctx context.Context, db *sql.DB, handler func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := handler(tx); err != nil {
		return err
	}

	tx.Commit()

	return nil
}
