package models

import (
	"math"
	"time"
)

type DateValue struct {
	Date  time.Time `json:"date"`
	Value float64   `json:"value"`
}

type WeatherGraphs struct {
	Temperature []DateValue `json:"temperature"`
	Humidity    []DateValue `json:"humidity"`
	Atmosphere  []DateValue `json:"atmosphere"`
	Co2         []DateValue `json:"co2"`
}

type Stat struct {
	Avg float64 `firestore:"avg"`
	Min float64 `firestore:"min"`
	Max float64 `firestore:"max"`
}

func DefaultStat() Stat {
	return Stat{0, math.Inf(1), math.Inf(-1)}
}

func SameStat(value float64) Stat {
	return Stat{value, value, value}
}

type WeatherSummary struct {
	Id          string    `firestore:"id"`
	Date        time.Time `firestore:"date"`
	Count       int       `firestore:"count"`
	Temperature Stat      `firestore:"temperature"`
	Humidity    Stat      `firestore:"humidity"`
	Atmosphere  Stat      `firestore:"atmosphere"`
	Co2         Stat      `firestore:"co2"`
}


