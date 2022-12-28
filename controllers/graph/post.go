package graph

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Nekodigi/Weather-Room/infrastructure"
	"github.com/Nekodigi/Weather-Room/models"
	"github.com/Nekodigi/Weather-Room/utils"
)

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
	doc, _ := client.Collection("weathers").Doc(req.Date.Format("060102")).Get(ctx)
	if doc == nil {
		log.Println("document not found => create new")
		var weatherGraphs models.WeatherGraphs
		utils.AppendGraphData(&weatherGraphs, req)
		wgJsonByte, _ := json.Marshal(weatherGraphs)
		log.Println(string(wgJsonByte))
		weatherSummary := models.WeatherSummary{
			Id:          req.Date.Format("060102"),
			Date:        req.Date,
			Cache:       string(wgJsonByte),
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
	doc.DataTo(&prevSummary)

	utils.UpdateWeatherSummary(&prevSummary, req)

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
