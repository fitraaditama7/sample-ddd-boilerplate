package migration

import (
	"database/sql"
	"ddd-boilerplate/pkg/logger"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	pgmigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func PostgresMigrate(db *sql.DB) error {
	log := logger.Logger

	driver, err := pgmigrate.WithInstance(db, &pgmigrate.Config{})
	if err != nil {
		log.Error(err.Error())
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migration",
		"postgres", driver)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	startVersion, _, err := m.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		log.Error(err.Error())
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange && err != migrate.ErrNilVersion && err != migrate.ErrLocked {
		log.Error("Migration is dirty, forcing rollback and retrying")
		endVersion, _, err := m.Version()
		if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
			log.Error(err.Error())
			return err
		}

		m.Force(int(endVersion) - 1)
		m.Steps((int(startVersion) + 1) - int(endVersion))
		m.Force(int(startVersion))
	}

	return nil
}
