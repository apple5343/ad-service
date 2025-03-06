package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	pgConnEnvName = "POSTGRES_DSN"
)

func main() {
	var migrationDir, connection string

	migrationDir = "./migrations"
	connection, _ = os.LookupEnv(pgConnEnvName)

	if migrationDir == "" || connection == "" {
		log.Println("POSTGRES_DSN is required")
		return
	}

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Printf("myapperror: %v", err)
		return
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Printf("myapperror: %v", err)
		return
	}

	fsrc, err := (&file.File{}).Open(fmt.Sprintf("file://%s", migrationDir))
	if err != nil {
		log.Printf("myapperror: %v", err)
		return
	}

	m, err := migrate.NewWithInstance("file", fsrc, "postgres", driver)
	if err != nil {
		log.Printf("myapperror: %v", err)
		return
	}
	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no changes")
			return
		}
		log.Printf("myapperror: %v", err)
		return
	}

	log.Println("success")
}
