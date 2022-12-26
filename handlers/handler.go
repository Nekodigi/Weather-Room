package handler

import (
	"net/http"
	"weather_room/controllers/graph"
	"weather_room/controllers/indicator"

	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	e.GET("/status", func(ctx *gin.Context) { ctx.String(http.StatusOK, "alive") })
	e.GET("/graph", graph.GET)
	e.POST("/graph", graph.POST)
	e.GET("/indicator", indicator.GET)
	e.POST("/indicator", indicator.POST)

}
