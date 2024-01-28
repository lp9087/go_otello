package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AddDashboardRoutes(route *gin.Engine) {
	dashboard := route.Group("/dashboard")
	dashboard.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "dashboard_pong")
	})
	dashboard.GET("/favourites_hotels/:id", GetFavouritesHotels)
	dashboard.POST("/favourites_hotels", GetFavouritesHotelsPost)
	dashboard.GET("/first_dashboard", GetFirstDashboard)
}

func GetFavouritesHotels(request *gin.Context) {
	request.String(http.StatusOK, request.Params.ByName("id"))
}

func GetFavouritesHotelsPost(request *gin.Context) {
	var Request struct {
		Message string
	}

	if err := request.BindJSON(&Request); err != nil {
		log.Fatal(err)
	}

	request.String(http.StatusOK, Request.Message)
}

func GetFirstDashboard(request *gin.Context) {

}
