package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lp9087/go_otello_dashboard_api/internal/entity"
	"github.com/lp9087/go_otello_dashboard_api/internal/usecase"
	"log/slog"
	"net/http"
)

type mostLoyalHotelRouter struct {
	useCase usecase.MostLoyalHotelsUseCase
	logger  *slog.Logger
}

func NewMostLoyalHotelsRoutes(handler *gin.RouterGroup, logger *slog.Logger, useCase usecase.MostLoyalHotelsUseCase) {
	r := &mostLoyalHotelRouter{useCase, logger}
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
		r.logger.Error("error", err, "http - v1 - mostLoyalHotels")
		errorResponse(c, http.StatusInternalServerError, "some API problems")
		return
	}
	response.Hotels = hotels

	c.JSON(http.StatusOK, response)
}
