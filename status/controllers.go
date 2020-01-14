package status

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type statusResponse struct {
	Message string `json:"message"`
}

// PingController Status application.
// @Summary Shows a very simple status report of the application.
// @Description Application status.
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /ping [get]
func pingController(c *gin.Context) {
	response := statusResponse{
		Message: "pong",
	}
	c.JSON(http.StatusOK, response)
}
