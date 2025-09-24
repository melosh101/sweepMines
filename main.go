package main

import (
	"flag"
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/melosh101/sweepMines/server"
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	hub := server.NewHub()
	go hub.Run()
	r := gin.Default()
	r.Use(static.Serve("/", GetFs()))
	r.SetTrustedProxies(nil)

	if isStatic {
		log.Println("Serving with embedded FS")
	} else {
		log.Println("Serving without embedded FS")
	}

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		api.GET("/ws", func(ctx *gin.Context) {
			w, r := ctx.Writer, ctx.Request
			log.Printf("%s %s", r.Method, r.URL.Path)
			server.ServeWs(hub, w, r)
		})
	}

	r.Run("0.0.0.0:8080")
}
