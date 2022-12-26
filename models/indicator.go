package models

import "time"

const (
	Good    = "good"
	Warning = "warning"
	Danger  = "danger"
)

type TGetIndicator struct {
	Date        time.Time `firestore:"date"`
	Temperature float64   `firestore:"temperature"`
	Humidity    float64   `firestore:"humidity"`
	Atmosphere  float64   `firestore:"atmosphere"`
	Co2         float64   `firestore:"co2"`
}

type WeatherIndicator struct {
	Date        time.Time `json:"date" firestore:"date"`
	Temperature Indicator `json:"temperature" firestore:"temperature"`
	Humidity    Indicator `json:"humidity" firestore:"humidity"`
	Atmosphere  Indicator `json:"atmosphere" firestore:"atmosphere"`
	Co2         Indicator `json:"co2" firestore:"co2"`
}

func SetIndicator(i *Indicator, v float64) Indicator {
	i.Current = v
	if i.GoodL <= v && v < i.GoodH {
		i.Status = Good
	} else if i.WarnL <= v && v < i.WarnH {
		i.Status = Warning
	} else {
		i.Status = Danger
	}
	return *i
}

type Indicator struct {
	Current float64 `json:"current" firestore:"current"`
	WarnL   float64 `json:"warnL" firestore:"warnL"`
	GoodL   float64 `json:"goodL" firestore:"goodL"`
	GoodH   float64 `json:"goodH" firestore:"goodH"`
	WarnH   float64 `json:"warnH" firestore:"warnH"`
	Status  string  `json:"status" firestore:"status"`
}
