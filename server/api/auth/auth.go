package auth

import "github.com/gin-gonic/gin"

func Serve(router *gin.RouterGroup) {
	router.POST("login", handleLogin)
	router.POST("register", handleRegister)
	router.GET("login", handleLogin)
}

func handleLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func handleRegister(c *gin.Context) {

}
