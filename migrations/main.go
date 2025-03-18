package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/mxpadidar/letsgo/internal/infra/conf"
)

func main() {
	configs := conf.NewConf()
	db, err := sql.Open("postgres", configs.PgDSN())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Create SQLite driver instance
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to create driver instance: %v", err)
	}

	// Create migration instance
	mgrtDir := "file://" + configs.RootDir + "/migrations/sql"
	mgrt, err := migrate.NewWithDatabaseInstance(mgrtDir, "postgres", driver)
	if err != nil {
		log.Fatalf("failed to create migration instance: %v", err)
	}

	// get the command from the arguments
	cmd := os.Args[len(os.Args)-1]
	switch cmd {
	case "up":
		err := mgrt.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to apply migration: %v", err)
		} else if err != nil && err == migrate.ErrNoChange {
			log.Println("db is up to date")
		} else {
			log.Println("migration applied successfully")
		}

	case "down":
		if err := mgrt.Down(); err != nil {
			log.Fatalf("failed to rollback migration: %v", err)
		} else {
			log.Println("migration rolled back successfully")
		}

	default:
		log.Fatalf("unknown command: %s", cmd)
	}

}
