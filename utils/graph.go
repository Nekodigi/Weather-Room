package utils

import (
	"math"
	"weather_room/models"
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

func UpdateStat(prev *models.Stat, count int, target float64) {
	prev.Avg = (prev.Avg*float64(count) + target) / float64(count+1)
	prev.Max = math.Max(prev.Max, target)
	prev.Min = math.Min(prev.Min, target)
}
