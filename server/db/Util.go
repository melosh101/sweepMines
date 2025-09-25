package db

import (
	"context"
	"embed"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed migrations
var embedFolder embed.FS

func Migrate(pool *pgxpool.Pool) {
	fmt.Println("running migrations...")
	migrationFiles, err := embedFolder.ReadDir("migrations")
	if err != nil {
		log.Fatal("failed to read migrations folder:", err)
	}

	for _, file := range migrationFiles {
		if file.IsDir() {
			continue
		}
		fmt.Println("running migration:", file.Name())
		migration, err := embedFolder.ReadFile("migrations/" + file.Name())
		if err != nil {
			log.Fatal("failed to read migration file:", err)
		}
		_, err = pool.Exec(context.Background(), string(migration))
		if err != nil {
			log.Fatal("failed to run migration:", err)
		}
	}

}
