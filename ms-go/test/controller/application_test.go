package controller

import (
	"encoding/json"
	"ms-go/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/likexian/gokit/assert"
)

func TestIndexHome(t *testing.T) {

	// Build our expected body
	body := gin.H{
		"message": "[ms-go] | Success",
		"status":  200,
	}

	// Grab our router
	router := router.SetupRouter()

	// Perform a GET request with that handler.
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert we encoded correctly,

	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the correct object response to a map
	response := struct {
		Data struct {
			Api struct {
				Status  int    `json:"status"`
				Message string `json:"message"`
			} `json:"api"`
		} `json:"data"`
	}{}

	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.Equal(t, body["message"], response.Data.Api.Message)

}
