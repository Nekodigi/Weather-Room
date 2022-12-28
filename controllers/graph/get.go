package graph

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Nekodigi/Weather-Room/infrastructure"
	"github.com/Nekodigi/Weather-Room/models"
	"github.com/Nekodigi/Weather-Room/utils"

	"github.com/gin-gonic/gin"
)

func GET(ctx *gin.Context) {
	log.Println("pastDays", ctx.Query("pastDays"))
	layout := "2006-01-02"
	target, err := time.Parse(layout, ctx.Query("target"))
	pastDays, _ := strconv.Atoi(ctx.Query("pastDays"))
	getAll, _ := strconv.ParseBool(ctx.Query("getAll"))

	if pastDays == 0 {
		log.Println("pastDays undefined => 1")
		pastDays = 1
	}
	if err != nil {
		log.Println("pastDays undefined => today")
		target = utils.JSTNow()
	}
	log.Println(target, err)
	if getAll {
		ReturnGraphAll(ctx, target)

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
			log.Println("Failed to get summary => bad request")
			ctx.String(http.StatusBadRequest, "Wrong pastDays\nFailed to get summary")
			return
		}
		var weatherSummary models.WeatherSummary
		doc.DataTo(&weatherSummary)
		weatherGraphs.Temperature = append(weatherGraphs.Temperature, models.DateValue{Date: weatherSummary.Date, Value: weatherSummary.Temperature.Avg})
		weatherGraphs.Humidity = append(weatherGraphs.Humidity, models.DateValue{Date: weatherSummary.Date, Value: weatherSummary.Humidity.Avg})
		weatherGraphs.Atmosphere = append(weatherGraphs.Atmosphere, models.DateValue{Date: weatherSummary.Date, Value: weatherSummary.Atmosphere.Avg})
		weatherGraphs.Co2 = append(weatherGraphs.Co2, models.DateValue{Date: weatherSummary.Date, Value: weatherSummary.Co2.Avg})
	}
	ctx.JSON(http.StatusOK, weatherGraphs)
}

func ReturnGraphAll(ctx *gin.Context, target time.Time) {
	client, _ := infrastructure.FirestoreInit(ctx)
	doc, err := client.Collection("weathers").Doc(target.Format("060102")).Get(ctx)
	if err != nil {
		log.Println("Failed to get cache => bad request")
		ctx.String(http.StatusBadRequest, "Wrong target\nFailed to get cache")
		return
	}
	var weatherSummary models.WeatherSummary
	doc.DataTo(&weatherSummary)

	var weatherGraphs models.WeatherGraphs
	json.Unmarshal([]byte(weatherSummary.Cache), &weatherGraphs)
	ctx.JSON(http.StatusOK, weatherGraphs)
}
