package handlers_test

import (
	"civi-backend-challenge/handlers"
	"civi-backend-challenge/models"
	"civi-backend-challenge/routers"
	"civi-backend-challenge/services"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	url = "/api/v1/points"
)

func getURLWithQueryParams(x, y, distance string) string {
	return fmt.Sprintf("%s?x=%s&y=%s&distance=%s", url, x, y, distance)
}

func TestHandlerGetPointsWithSuccess(t *testing.T) {
	data := []*models.Points{
		{X: 5, Y: 2},
		{X: 4, Y: 2},
		{X: 2, Y: 1},
		{X: 3, Y: 1},
		{X: 13, Y: 11},
	}

	s := services.NewPointService(data)

	h := handlers.New(s)
	r := routers.InitRouter(h)

	url := getURLWithQueryParams("3", "2", "10")
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response []map[string]int
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Len(t, response, 4)
}
