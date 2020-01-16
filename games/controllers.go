package games

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateGameController Creates a new game.
// @Summary Create a new game instance with dummy configuration.
// @Description New game instance.
// @Accept  json
// @Produce  json
// @Success 201 {object} Game
// @Router /games [post]
func createGameController(c *gin.Context) {
	game := New()
	c.JSON(http.StatusCreated, game)
}
