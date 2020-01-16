package main

import (
	"github.com/ariel17/minesweeper-API/config"
	_ "github.com/ariel17/minesweeper-API/docs"
	"github.com/ariel17/minesweeper-API/games"
	"github.com/ariel17/minesweeper-API/status"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func main() {
	loadRoutes(gin.Default()).Run()
}

func loadRoutes(r *gin.Engine) *gin.Engine {
	status.LoadRoutes(r)
	games.LoadRoutes(r)
	url := ginSwagger.URL(config.PublicURL + "/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
