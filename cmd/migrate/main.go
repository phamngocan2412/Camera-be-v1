package main

import (
	"flag"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL required")
	}

	var path string
	flag.StringVar(&path, "path", "file://../../migrations", "migration path")
	flag.Parse()

	command := flag.Arg(0)

	m, err := migrate.New(path, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	switch command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Up success")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Down success")
	case "version":
		v, dirty, _ := m.Version()
		log.Printf("Version: %d (dirty: %v)", v, dirty)
	default:
		log.Println("Usage: go run ./cmd/migrate [up|down|version]")
	}
}
