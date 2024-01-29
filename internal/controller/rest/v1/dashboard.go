package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lp9087/go_otello_dashboard_api/internal/entity"
	"github.com/lp9087/go_otello_dashboard_api/internal/usecase"
	"net/http"
)

type dashboardRoutes struct {
	t usecase.FirstDashboard
}

func newDashboardRoutes(handler *gin.RouterGroup, t usecase.FirstDashboard) {
	r := &dashboardRoutes{t}

	h := handler.Group("/dashboard")
	{
		h.GET("/mostLoyalHotels", r.mostLoyalHotels)
	}
}

type loyalHotelsResponse struct {
	History []entity.FirstDashboard `json:"mostLoyalHotels"`
}

// @Summary     Show mostLoyalHotels
// @Description Show all dashboards mostLoyalHotels
// @ID          mostLoyalHotels
// @Tags  	    dashboard
// @Accept      json
// @Produce     json
// @Success     200 {object} loyalHotelsResponse
// @Failure     500 {object} response
// @Router      /translation/mostLoyalHotels [get]
func (r *dashboardRoutes) mostLoyalHotels(c *gin.Context) {
	hotels, err := r.t.Get(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "some API problems")
		return
	}

	c.JSON(http.StatusOK, loyalHotelsResponse{hotels})
}
