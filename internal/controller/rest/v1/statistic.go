package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lp9087/go_otello_dashboard_api/internal/entity"
	"github.com/lp9087/go_otello_dashboard_api/internal/usecase"
	"log/slog"
	"net/http"
	"time"
)

type statisticRouter struct {
	useCase usecase.HotelStatisticUseCase
	logger  *slog.Logger
}

func NewStatisticRoutes(handler *gin.RouterGroup, logger *slog.Logger, useCase usecase.HotelStatisticUseCase) {
	r := &statisticRouter{useCase, logger}
	{
		handler.GET("/registered_hotels", r.hotelsStatistic)
	}
}

type hotelStatisticResponse struct {
	Statistic entity.HotelStatistic `json:"HotelStatistic"`
}

// @Summary		Show hotelsStatistic
// @Description	Show amount of hotelsStatistic
// @ID				hotelsStatistic
// @Tags			dashboard
// @Accept			json
// @Produce		json
// @Success		200	{object}	hotelStatisticResponse
// @Failure		500	{object}	response
// @Router			/dashboard/hotel_statistic [get]
func (r *statisticRouter) hotelsStatistic(c *gin.Context) {
	var response hotelStatisticResponse
	strDateFrom, ok := c.GetQuery("date_from")
	strDateTo, ok := c.GetQuery("date_to")
	if ok != true {
		r.logger.Error("error", "http - v1 - hotelsStatistic")
		errorResponse(c, http.StatusBadRequest, "not enough params")
	}
	var err error
	dateFrom, err := time.Parse("2006-01-02", strDateFrom)
	dateTo, err := time.Parse("2006-01-02", strDateTo)
	if err != nil {
		r.logger.Error("error", "http - v1 - hotelsStatistic")
		errorResponse(c, http.StatusBadRequest, "cannot convert params to datetime")
		return
	}

	statistics, err := r.useCase.Get(c.Request.Context(), dateFrom, dateTo)
	if err != nil {
		r.logger.Error("error", err, "http - v1 - hotelsStatistic")
		errorResponse(c, http.StatusInternalServerError, "some API problems")
		return
	}
	response.Statistic = *statistics

	c.JSON(http.StatusOK, response)
}
