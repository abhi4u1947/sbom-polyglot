package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
}

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", homeHandler)

	return router
}

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, response{
		Message: "Generating the SBOM!",
	})
}
