package models

import "time"

type TWeatherPost struct {
	Date        time.Time `json:"date"`
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	Atmosphere  float64   `json:"atmosphere"`
	Co2         float64   `json:"co2"`
}

type WeatherData struct {
	Id          string    `firestore:"id"`
	Date        time.Time `firestore:"date"`
	Temperature float64   `firestore:"temperature"`
	Humidity    float64   `firestore:"humidity"`
	Atmosphere  float64   `firestore:"atmosphere"`
	Co2         float64   `firestore:"co2"`
}
