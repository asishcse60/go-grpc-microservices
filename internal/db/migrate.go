package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func (s *Store) Migrate()error {
	driver,err := postgres.WithInstance(s.db.DB, &postgres.Config{})
	if err != nil{
		fmt.Println(err)
		return err
	}
	m,err:=migrate.NewWithDatabaseInstance("file:///migrations", "postgres", driver)
	if err != nil{
		fmt.Println(err)
		return err
	}
	m.Steps(2)
	return nil
}
