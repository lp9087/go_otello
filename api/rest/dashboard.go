package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddDashboardRoutes(route *gin.Engine) {
	dashboard := route.Group("/dashboard")
	dashboard.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "dashboard_pong")
	})
}
