package api

import (
	"github.com/gin-gonic/gin"
	"github.com/melosh101/sweepMines/server/api/auth"
)

func Serve(router *gin.RouterGroup) {

	router.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authRouter := router.Group("auth")
	auth.Serve(authRouter)
}
