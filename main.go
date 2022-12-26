package main

import (
	"context"
	"math"
	"os"
	"time"
	"weather_room/controllers/graph"
	handler "weather_room/handlers"
	"weather_room/models"
	"weather_room/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// for i, v := range os.Args {
	// 	log.Printf("args[%d] -> %s\n", i, v)
	// }
	if len(os.Args) == 2 && os.Args[1] == "test" {
		test()
	} else {
		engine := gin.Default()
		handler.Router(engine)
		engine.Run(":3000")
	}
}

type GoStruct struct {
	A int
	B string
}

type JSON struct {
	V1 string
	V2 int
	V3 int
}

type NestJSON struct {
	J1 JSON
	Id string
}

func test() {
	for min := 0.0; min < 10*6*12; min += 10 * 6 {
		theta := 2.0 * math.Pi * min / (10 * 6 * 24)
		req := models.TWeatherPost{
			Date:        time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, int(min), 0, 0, utils.JST()),
			Temperature: math.Sin(theta*2)*10 + 10,
			Humidity:    math.Sin(theta*2+math.Pi/2)*50 + 100,
			Atmosphere:  math.Sin(theta*2+math.Pi)*500 + 500,
			Co2:         math.Sin(theta*2+math.Pi/2*3)*5000 + 5000,
		}
		graph.AddData(context.Background(), req)
	}

}
