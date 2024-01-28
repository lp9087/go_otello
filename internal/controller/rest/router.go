package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRoutes() *gin.Engine {
	router := SetupRouter()
	AddDashboardRoutes(router)
	return router
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}
