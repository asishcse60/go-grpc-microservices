package db

import (
	"fmt"
	"github.com/asishcse60/go-grpc-microservices/internal/rocket"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Store struct {
	db *sqlx.DB
}

func NewDatabase() (Store, error) {
	fmt.Println("Setting up new database connection")
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSL_MODE")

	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUsername, dbName, dbPassword, sslMode)
	fmt.Println(dbUsername)
	fmt.Println(dbPassword)
	fmt.Println(dbHost)
	fmt.Println(dbName)
	fmt.Println(dbPort)
	db, err := sqlx.Connect("postgres", connectString)
	if err != nil {
		fmt.Println(err)
		return Store{}, err
	}
	return Store{db: db}, nil
}

// GetRocketByID - returns a rocket from the database by a given ID
func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

// InsertRocket - inserts a new rocket into the database
func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

// DeleteRocket - deletes a rocket from the database by it's ID
func (s Store) DeleteRocket(id string) error {
	return nil
}
