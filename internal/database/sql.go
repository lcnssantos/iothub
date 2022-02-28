package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lcnssantos/iothub/cmd/publicApi/config"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() (*sql.DB, error) {
	configuration := config.Get()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", configuration.DB_USER, configuration.DB_PASS, configuration.DB_HOST, configuration.DB_PORT, configuration.DB_NAME)

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

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func GetTransaction(ctx context.Context, db *sql.DB) (*sql.Tx, error) {
	tx, err := db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	return tx, nil
}
