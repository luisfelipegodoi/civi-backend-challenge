package handlers

import (
	"civi-backend-challenge/models"
	"civi-backend-challenge/services"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler interface {
	GetPoints(c *gin.Context)
}

type handler struct {
	pointsService services.PointsService
}

func New(s services.PointsService) *handler {
	return &handler{
		pointsService: s,
	}
}

// GetPoints godoc
// @Summary      Get points
// @Description  get points by params
// @Tags         points
// @Accept       json
// @Produce      json
// @Param        x   query      int  true  "x"
// @Param        y   query      int  true  "y"
// @Param        distance   query      int  true  "distance"
// @Success      200  {object}  models.Points
// @Failure      400  {int}  http.StatusBadRequest
// @Router       /api/v1/points/{x} [get]
func (h *handler) GetPoints(c *gin.Context) {
	queryParam := c.Request.URL.Query()

	x, ok, err := isInteger(queryParam.Get("x"))
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{
			StatusCode: http.StatusBadRequest,
			Message:    models.XParamInvalidOrNullError.Error(),
			Details:    err.Error(),
		})
		return
	}

	y, ok, err := isInteger(queryParam.Get("y"))
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{
			StatusCode: http.StatusBadRequest,
			Message:    models.YParamInvalidOrNullError.Error(),
			Details:    err.Error(),
		})
		return
	}

	distance, ok, err := isInteger(queryParam.Get("distance"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{
			StatusCode: http.StatusBadRequest,
			Message:    models.DistanceParamInvalidOrNullError.Error(),
			Details:    err.Error(),
		})
		return
	}

	points := h.pointsService.GetPoints(x, y, distance)

	if points == nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"Message": "there's no result for show. the parameters don't return any results",
		})
		return
	}

	c.JSON(http.StatusOK, &points)
}

func isInteger(param string) (int, bool, error) {
	var number int

	if param == "" {
		return number, false, errors.New("param is wrong, null or empty")
	}

	number, err := strconv.Atoi(param)
	if err != nil {
		return number, false, err
	}

	return number, true, nil
}
