package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", WelcomeHandler)
	router.POST("/register", RegisterHandler)
}

func WelcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Bem-vindo ao servidor Gin!",
	})
}

func RegisterHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Registro realizado com sucesso!",
	})
}
