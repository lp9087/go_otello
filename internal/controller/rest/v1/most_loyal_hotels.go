package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lp9087/go_otello_dashboard_api/internal/entity"
	"github.com/lp9087/go_otello_dashboard_api/internal/usecase"
	"log"
	"net/http"
)

type mostLoyalHotelRouter struct {
	useCase usecase.MostLoyalHotelsUseCase
}

func NewMostLoyalHotelsRoutes(handler *gin.RouterGroup, useCase usecase.MostLoyalHotelsUseCase) {
	r := &mostLoyalHotelRouter{useCase}
	{
		handler.GET("/mostLoyalHotels", r.mostLoyalHotels)
	}
}

type loyalHotelsResponse struct {
	Hotels []entity.MostLoyalHotels `json:"mostLoyalHotels"`
}

// @Summary		Show mostLoyalHotels
// @Description	Show all dashboards mostLoyalHotels
// @ID				mostLoyalHotels
// @Tags			dashboard
// @Accept			json
// @Produce		json
// @Success		200	{object}	loyalHotelsResponse
// @Failure		500	{object}	response
// @Router			/dashboard/mostLoyalHotels [get]
func (r *mostLoyalHotelRouter) mostLoyalHotels(c *gin.Context) {
	var response loyalHotelsResponse
	hotels, err := r.useCase.Get(c.Request.Context())
	if err != nil {
		//TODO add logger slog
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "some API problems")
		return
	}
	response.Hotels = hotels

	c.JSON(http.StatusOK, response)
}
