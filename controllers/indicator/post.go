package indicator

import (
	"net/http"
	"time"

	"github.com/Nekodigi/Weather-Room/controllers/graph"
	"github.com/Nekodigi/Weather-Room/infrastructure"
	"github.com/Nekodigi/Weather-Room/models"
	"github.com/Nekodigi/Weather-Room/utils"

	"github.com/gin-gonic/gin"
)

func POST(ctx *gin.Context) {
	//setup data
	var req models.TWeatherPost
	ctx.BindJSON(&req)
	if req.Date.IsZero() {
		req.Date = utils.JSTNow()
	}

	needToUpdateGraph := UpdateIndicator(ctx, req)
	//doc date
	if needToUpdateGraph {
		//snap date
		req.Date = time.Date(req.Date.Year(), req.Date.Month(), req.Date.Day(), req.Date.Hour(), (req.Date.Minute()/models.UpdateEachMin)*models.UpdateEachMin, 0, 0, utils.JST())
		graph.AddData(ctx, req)
		ctx.String(http.StatusOK, "graph updated")
	} else {
		ctx.String(http.StatusOK, "indicator updated")
	}
}

// return need for graph update
func UpdateIndicator(ctx *gin.Context, req models.TWeatherPost) bool {
	//update indicator
	client, err := infrastructure.FirestoreInit(ctx)
	data := models.WeatherData{
		Id:          "latest",
		Date:        req.Date,
		Temperature: req.Temperature,
		Humidity:    req.Humidity,
		Atmosphere:  req.Atmosphere,
		Co2:         req.Co2,
	}
	if err != nil {
		ctx.String(http.StatusInternalServerError, "firebase init error")
	}
	doc, _ := client.Collection("weathers").Doc("latest").Get(ctx) //if not exist, doc will be nil

	var prevData models.WeatherData
	doc.DataTo(&prevData)

	needToUpdateGraph := false
	if doc != nil {
		needToUpdateGraph = prevData.Date.Minute()/models.UpdateEachMin != utils.JSTNow().Minute()/models.UpdateEachMin
	} else {
		needToUpdateGraph = true
	}
	client.Collection("weathers").Doc("latest").Set(ctx, data)
	return needToUpdateGraph
}
