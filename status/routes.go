package status

import "github.com/gin-gonic/gin"

// LoadRoutes configures controllers for the given router.
func LoadRoutes(r *gin.Engine) {
	r.GET("/ping", pingController)
}
