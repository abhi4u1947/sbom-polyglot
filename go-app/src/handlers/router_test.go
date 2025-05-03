package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	// Create a new router
	router := NewRouter()

	// Create a test server
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Test the home endpoint
	t.Run("GET /", func(t *testing.T) {
		// Make a GET request to the home endpoint
		resp, err := http.Get(ts.URL + "/")
		assert.NoError(t, err)
		defer resp.Body.Close()

		// Check the status code
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Check the content type
		assert.Equal(t, "application/json; charset=utf-8", resp.Header.Get("Content-Type"))
	})
}

func TestHomeHandler(t *testing.T) {
	// Create a test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call the handler
	homeHandler(c)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	expected := `{"message":"Generating the SBOM!"}`
	assert.Equal(t, expected, w.Body.String())
}
