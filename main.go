package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/melosh101/sweepMines/server"
	"github.com/melosh101/sweepMines/server/api"
	"github.com/melosh101/sweepMines/server/db"
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	hub := server.NewHub()
	go hub.Run()
	r := gin.Default()
	r.Use(static.Serve("/", GetFs()))
	r.SetTrustedProxies(nil)

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	db.Migrate(pool)

	if isStatic {
		log.Println("Serving with embedded FS")
	} else {
		log.Println("Serving without embedded FS")
	}
	apiRouter := r.Group("/api")
	api.Serve(apiRouter)

	r.Run("0.0.0.0:8080")
}
