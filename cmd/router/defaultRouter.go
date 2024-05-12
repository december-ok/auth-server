package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addDefaultRoutes(rg *gin.RouterGroup) {
	rg.POST("/join", func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{
				"message": "join",
			})
	})
}
