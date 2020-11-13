package main

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})
	log.SetLevel(log.DebugLevel)

	log.Info("Hello, i have nothing to show you yet :(")

	dbUrl := "We don't have one yet..."

	db, err := sqlx.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)

	if err := runMigrations(db); err != nil {
		log.Fatal(err)
	}

	return
}

func runMigrations(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "Could not create the driver")
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return errors.Wrap(err, "Could not initiate the migrations")
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return errors.Wrap(err, "Could not migrate")
		}
	}

	return nil
}
