package status

import "github.com/gin-gonic/gin"

// LoadRoutes configures controllers for the given router.
func LoadRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
