package main

import (
	"github.com/eron97/testes-mocks.git/teste2/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}
