package graph

import (
	"context"
	"log"
	"net/http"
	"time"
	"weather_room/infrastructure"
	"weather_room/models"
	"weather_room/utils"

	"github.com/gin-gonic/gin"
)

func POST(ctx *gin.Context) {

	var req models.TWeatherPost
	ctx.BindJSON(&req)
	if req.Date.IsZero() {
		req.Date = utils.JSTNow()
	}
	AddData(ctx, req)
	ctx.String(http.StatusOK, "updated")
	// /weather/YYMM/datas/hhmm
}

func AddData(ctx context.Context, req models.TWeatherPost) {
	client, _ := infrastructure.FirestoreInit(ctx)

	data := models.WeatherData{
		Id:          req.Date.Format("1504"),
		Date:        req.Date,
		Temperature: req.Temperature,
		Humidity:    req.Humidity,
		Atmosphere:  req.Atmosphere,
		Co2:         req.Co2,
	}

	UpdateStat(ctx, req) //* should be called
	//add new data
	_, err := client.Collection("weathers").Doc(req.Date.Format("060102")).Collection("datas").Doc(req.Date.Format("1504")).Set(ctx, data)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Println("An error has occurred:", err)
	}

}

func UpdateStat(ctx context.Context, req models.TWeatherPost) {
	client, _ := infrastructure.FirestoreInit(ctx)

	var prevSummary models.WeatherSummary
	ref, err := client.Collection("weathers").Doc(req.Date.Format("060102")).Get(ctx)
	if err != nil {
		log.Println("document not found => create new")
		weatherSummary := models.WeatherSummary{
			Id:          req.Date.Format("060102"),
			Date:        time.Date(req.Date.Year(), req.Date.Month(), req.Date.Day(), 0, 0, 0, 0, utils.JST()),
			Temperature: models.SameStat(req.Temperature),
			Humidity:    models.SameStat(req.Humidity),
			Atmosphere:  models.SameStat(req.Atmosphere),
			Co2:         models.SameStat(req.Co2),
			Count:       1,
		}

		_, err := client.Collection("weathers").Doc(req.Date.Format("060102")).Set(ctx, weatherSummary)
		if err != nil {
			log.Println("cannot create document: ", err)
		}
		return
	}
	ref.DataTo(&prevSummary)

	utils.UpdateStat(&prevSummary.Temperature, prevSummary.Count, req.Temperature)
	utils.UpdateStat(&prevSummary.Humidity, prevSummary.Count, req.Humidity)
	utils.UpdateStat(&prevSummary.Atmosphere, prevSummary.Count, req.Atmosphere)
	utils.UpdateStat(&prevSummary.Co2, prevSummary.Count, req.Co2)
	prevSummary.Count++

	{
		_, err := client.Collection("weathers").Doc(req.Date.Format("060102")).Set(ctx, prevSummary)
		if err != nil {
			log.Println("cannot create document: ", err)
		}
	}
	// weatherDatas := []models.WeatherData{}
	// iter := client.Collection("weather").Doc(req.Date.Format("060102")).Collection("datas").Documents(ctx)
	// for {
	// 	doc, err := iter.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatalf("Failed to iterate: %v", err)
	// 	}
	// 	var weatherData models.WeatherData
	// 	doc.DataTo(&weatherData)
	// 	weatherDatas = append(weatherDatas, weatherData)
	// }
	// weatherGraphs := WeatherDataToGraph(weatherDatas)
	// weatherSummary.Temperature = utils.GraphToStat(weatherGraphs.Temperature)
	// weatherSummary.Humidity = utils.GraphToStat(weatherGraphs.Humidity)
	// weatherSummary.Atmosphere = utils.GraphToStat(weatherGraphs.Atmosphere)
	// weatherSummary.Co2 = utils.GraphToStat(weatherGraphs.Co2)

	//update stat
	// _, err := client.Collection("weather").Doc(req.Date.Format("060102")).Set(ctx, weatherSummary)
	// if err != nil {
	// 	// Handle any errors in an appropriate way, such as returning them.
	// 	log.Println("An error has occurred: ", err)
	// }
}
