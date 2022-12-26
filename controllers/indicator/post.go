package indicator

import (
	"net/http"
	"weather_room/infrastructure"
	"weather_room/models"
	"weather_room/utils"

	"github.com/gin-gonic/gin"
)

func POST(ctx *gin.Context) {
	var req models.TWeatherPost
	ctx.BindJSON(&req)

	client, _ := infrastructure.FirestoreInit(ctx)
	data := models.WeatherData{
		Id:          "latest",
		Date:        utils.JSTNow(),
		Temperature: req.Temperature,
		Humidity:    req.Humidity,
		Atmosphere:  req.Atmosphere,
		Co2:         req.Co2,
	}
	client.Collection("weathers").Doc("latest").Set(ctx, data)
	ctx.String(http.StatusOK, "updated")
}
