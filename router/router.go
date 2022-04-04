package router

import (
	"starbuy/router/routes"

	"github.com/gin-gonic/gin"
)

// Build vai retornar um router com as rotas configuradas
func Build() *gin.Engine {
	return routes.Configure(gin.Default())
}
