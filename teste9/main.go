package main

import (
	"github.com/eron97/testes-mocks.git/teste9/controller/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	router.Run(":8080")
}
