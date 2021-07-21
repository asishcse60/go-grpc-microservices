package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
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
	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("no change made by migrations")
		} else {
			return err
		}
	}
	return nil
}
