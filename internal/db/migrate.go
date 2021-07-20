package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"os"
)

func (s *Store) Migrate() error {
	driver, err := postgres.WithInstance(s.db.DB, &postgres.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	dbName := os.Getenv("DB_NAME")
	m, err := migrate.NewWithDatabaseInstance("file://migrations", dbName, driver)
	if err != nil {
		fmt.Println(err)
		return err
	}
	m.Steps(2)
	return nil
}
