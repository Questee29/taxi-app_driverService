package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upDrivers, downDrivers)
}

func upDrivers(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS drivers(
		"id" SERIAL PRIMARY KEY,
		"name" text,
		"phone" text,
		"email" text,
		"password" text,
		"taxi_type" text,
		"status" text,
		"rating" real);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downDrivers(tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS drivers;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
