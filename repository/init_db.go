package repository

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var (
	DB       *Queries
	dbClient *sql.DB
)

func ConnectDB() error {
	connStr := "postgresql://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable"
	var err error
	dbClient, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err = dbClient.Ping(); err != nil {
		return err
	}

	DB = New(dbClient)

	return nil
}

func RunMigrations() error {
	driver, err := postgres.WithInstance(dbClient, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://repository/migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		// migration down?
		return err
	}

	return nil
}
