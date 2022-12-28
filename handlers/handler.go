package handler

import (
	"net/http"

	"github.com/Nekodigi/Weather-Room/controllers/graph"
	"github.com/Nekodigi/Weather-Room/controllers/indicator"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Router(e *gin.Engine) {
	e.Use(CORSMiddleware())
	e.GET("/status", func(ctx *gin.Context) { ctx.String(http.StatusOK, "alive") })
	e.GET("/graph", graph.GET)
	e.POST("/graph", indicator.POST)
	e.GET("/indicator", indicator.GET)
	e.POST("/indicator", indicator.POST)

}
