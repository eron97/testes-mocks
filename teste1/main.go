// main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cria uma instância do roteador Gin
	router := gin.Default()

	// Define uma rota GET
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Bem-vindo ao servidor Gin!",
		})
	})

	// Inicia o servidor na porta 8080
	router.Run(":8080")
}
