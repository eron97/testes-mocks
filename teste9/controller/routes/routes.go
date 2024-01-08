package routes

import (
	"github.com/eron97/testes-mocks.git/teste9/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/createUser", controller.CreateUser)
}
