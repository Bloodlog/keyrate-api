package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"key-rate-api/src/Routes"

	"github.com/stretchr/testify/assert"
)

func TestKeyrateRoute(t *testing.T) {
	router := Routes.SetUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/keyrate", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `[ { "date": "2022-01-24T00:00:00+03:00", "rate": "8.50" } ]`, w.Body.String())
}
