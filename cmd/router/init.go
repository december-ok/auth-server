package router

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const PORT = 8080

func RunServer() {
	webServer := gin.Default()

	getRoutes(webServer)
	webServer.Run(":" + strconv.Itoa(PORT))
}

func getRoutes(webServer *gin.Engine) {
	defaultRouter := webServer.Group("/")
	addDefaultRoutes(defaultRouter)
}
