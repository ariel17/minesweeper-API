package main

import (
	"github.com/ariel17/minesweeper-API/status"
	"github.com/gin-gonic/gin"
)

func main() {
	loadRoutes(gin.Default()).Run()
}

func loadRoutes(r *gin.Engine) *gin.Engine {
	status.LoadRoutes(r)
	return r
}
