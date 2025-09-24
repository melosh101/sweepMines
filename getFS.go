//go:build !dev

package main

import (
	"embed"
	"log"

	"github.com/gin-contrib/static"
)

//go:embed client/dist
var embedFS embed.FS

const isStatic = true

func GetFs() static.ServeFileSystem {

	embedFolder, err := static.EmbedFolder(embedFS, "client/dist")
	if err != nil {
		log.Fatal("failed to get embedded static filesystem")
	}
	return embedFolder
}
