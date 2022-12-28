package graph

import "github.com/Nekodigi/Weather-Room/models"

func WeatherDataToGraph(weatherDatas []models.WeatherData) models.WeatherGraphs {
	var weatherGraphs models.WeatherGraphs
	for _, weatherData := range weatherDatas {
		dateTemperature := models.DateValue{weatherData.Date, weatherData.Temperature}
		weatherGraphs.Temperature = append(weatherGraphs.Temperature, dateTemperature)
		dateHumidity := models.DateValue{weatherData.Date, weatherData.Humidity}
		weatherGraphs.Humidity = append(weatherGraphs.Humidity, dateHumidity)
		dateAtmosphere := models.DateValue{weatherData.Date, weatherData.Atmosphere}
		weatherGraphs.Atmosphere = append(weatherGraphs.Atmosphere, dateAtmosphere)
		dateCo2 := models.DateValue{weatherData.Date, weatherData.Co2}
		weatherGraphs.Co2 = append(weatherGraphs.Co2, dateCo2)
	}
	return weatherGraphs
}
