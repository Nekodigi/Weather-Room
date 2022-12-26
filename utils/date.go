package utils

import (
	"log"
	"time"
)

func JSTNow() time.Time {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalln("invalid time zone", err)
	}
	return time.Now().In(jst)
}
func JST() *time.Location {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalln("invalid time zone", err)
	}
	return jst
}
