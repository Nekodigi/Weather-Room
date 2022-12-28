package utils

import (
	"encoding/json"
	"math"

	"github.com/Nekodigi/Weather-Room/models"
)

func GraphToStat(g []models.DateValue) models.Stat {
	var Min float64
	var Avg float64
	var Max float64

	for _, v := range g {
		Max = math.Max(v.Value, Max)
		Min = math.Min(v.Value, Min)
		Avg += v.Value
	}
	Avg /= float64(len(g))
	return models.Stat{Avg: Avg, Min: Min, Max: Max}
}

func AppendGraphData(target *models.WeatherGraphs, req models.TWeatherPost) {
	target.Temperature = append(target.Temperature, models.DateValue{Date: req.Date, Value: req.Temperature})
	target.Humidity = append(target.Humidity, models.DateValue{Date: req.Date, Value: req.Humidity})
	target.Atmosphere = append(target.Atmosphere, models.DateValue{Date: req.Date, Value: req.Atmosphere})
	target.Co2 = append(target.Co2, models.DateValue{Date: req.Date, Value: req.Co2})
}

func UpdateWeatherSummary(weatherSummary *models.WeatherSummary, req models.TWeatherPost) {
	UpdateStat(&weatherSummary.Temperature, weatherSummary.Count, req.Temperature)
	UpdateStat(&weatherSummary.Humidity, weatherSummary.Count, req.Humidity)
	UpdateStat(&weatherSummary.Atmosphere, weatherSummary.Count, req.Atmosphere)
	UpdateStat(&weatherSummary.Co2, weatherSummary.Count, req.Co2)
	var weatherGraphs models.WeatherGraphs
	json.Unmarshal([]byte(weatherSummary.Cache), &weatherGraphs)
	AppendGraphData(&weatherGraphs, req)
	wgJsonByte, _ := json.Marshal(weatherGraphs)
	weatherSummary.Cache = string(wgJsonByte)
	weatherSummary.Count++
}

func UpdateStat(prev *models.Stat, count int, target float64) {
	prev.Avg = (prev.Avg*float64(count) + target) / float64(count+1)
	prev.Max = math.Max(prev.Max, target)
	prev.Min = math.Min(prev.Min, target)
}
