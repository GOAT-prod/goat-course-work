package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func ConnectPostgres(connectionString string) *sqlx.DB {
	return sqlx.MustConnect("postgres", connectionString)
}

func RunMigrations(database *sqlx.DB) error {
	goose.SetVerbose(false)
	goose.SetTableName("client_service_db_versions")

	return goose.Up(database.DB, "./database/migrations")
}
