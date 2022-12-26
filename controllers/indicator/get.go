package indicator

import (
	"log"
	"net/http"
	"weather_room/infrastructure"
	"weather_room/models"

	"github.com/gin-gonic/gin"
)

func GET(ctx *gin.Context) {
	client, _ := infrastructure.FirestoreInit(ctx)

	doc, err := client.Collection("weathers").Doc("latest").Get(ctx)
	if err != nil {
		log.Fatalln("cannot get latest")
	}
	var weatherData models.TGetIndicator
	doc.DataTo(&weatherData)

	doc2, err2 := client.Collection("consts").Doc("weatherCriteria").Get(ctx)
	if err2 != nil {
		log.Fatalln("cannot get criteria")
	}
	var weatherCriteria models.WeatherIndicator
	doc2.DataTo(&weatherCriteria)

	weatherIndicator := models.WeatherIndicator{
		Date:        weatherData.Date,
		Temperature: models.SetIndicator(&weatherCriteria.Temperature, weatherData.Temperature),
		Humidity:    models.SetIndicator(&weatherCriteria.Humidity, weatherData.Humidity),
		Atmosphere:  models.SetIndicator(&weatherCriteria.Atmosphere, weatherData.Atmosphere),
		Co2:         models.SetIndicator(&weatherCriteria.Co2, weatherData.Co2),
	}
	ctx.JSON(http.StatusOK, weatherIndicator)
}
