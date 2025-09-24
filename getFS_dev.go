//go:build dev

package main

import "github.com/gin-contrib/static"

const isStatic = false

func GetFs() static.ServeFileSystem {
	return static.LocalFile("./client/dist", true)
}
