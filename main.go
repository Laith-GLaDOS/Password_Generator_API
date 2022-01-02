package main

import (
	"passgen_api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	server.Use(LoggerMiddleware)
	server.GET("/", routes.Root)
	server.GET("/api", routes.Index)
	server.GET("/api/docs", routes.Docs)
	server.POST("/api", routes.API)
	server.Run(os.Getenv("passgen_api_address"))
}
