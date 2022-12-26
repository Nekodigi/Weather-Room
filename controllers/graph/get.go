package graph

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"weather_room/infrastructure"
	"weather_room/models"
	"weather_room/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func GET(ctx *gin.Context) {
	log.Println("pastDays", ctx.Query("pastDays"))
	pastDays, _ := strconv.Atoi(ctx.Query("pastDays"))
	getAll, _ := strconv.ParseBool(ctx.Query("getAll"))

	if pastDays == 0 {
		log.Println("pastDays undefined => 1")
		pastDays = 1
	}
	if getAll {
		ReturnGraphAll(ctx, pastDays)

	} else {
		ReturnGraphAvg(ctx, pastDays)
	}

}

func ReturnGraphAvg(ctx *gin.Context, pastDays int) {
	client, _ := infrastructure.FirestoreInit(ctx)
	date := utils.JSTNow().AddDate(0, 0, -pastDays)
	weatherGraphs := models.WeatherGraphs{}

	for i := 0; i < pastDays; i++ {
		date = date.AddDate(0, 0, 1)
		doc, err := client.Collection("weathers").Doc(date.Format("060102")).Get(ctx)
		if err != nil {
			log.Fatalf("Failed to get avg: %v", err)
		}
		var weatherSummary models.WeatherSummary
		doc.DataTo(&weatherSummary)
		weatherGraphs.Temperature = append(weatherGraphs.Temperature, models.DateValue{weatherSummary.Date, weatherSummary.Temperature.Avg})
		weatherGraphs.Humidity = append(weatherGraphs.Humidity, models.DateValue{weatherSummary.Date, weatherSummary.Humidity.Avg})
		weatherGraphs.Atmosphere = append(weatherGraphs.Atmosphere, models.DateValue{weatherSummary.Date, weatherSummary.Atmosphere.Avg})
		weatherGraphs.Co2 = append(weatherGraphs.Co2, models.DateValue{weatherSummary.Date, weatherSummary.Co2.Avg})
	}
	ctx.JSON(http.StatusOK, weatherGraphs)
}

func ReturnGraphAll(ctx *gin.Context, pastDays int) {
	client, _ := infrastructure.FirestoreInit(ctx)
	date := utils.JSTNow().AddDate(0, 0, -pastDays)
	weatherDatas := []models.WeatherData{}
	for i := 0; i < pastDays; i++ {
		iter := client.Collection("weathers").Doc(date.Format("060102")).Collection("datas").Documents(ctx)
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("Failed to iterate: %v", err)
			}
			var weatherData models.WeatherData
			doc.DataTo(&weatherData)
			weatherDatas = append(weatherDatas, weatherData)
		}
	}
	weatherGraphs := WeatherDataToGraph(weatherDatas)
	fmt.Println(WeatherDataToGraph(weatherDatas))

	ctx.JSON(http.StatusOK, weatherGraphs)
}
