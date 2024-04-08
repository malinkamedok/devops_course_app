package usecase

import (
	"devops_course_app/internal/entity/weather"
	"fmt"
	"slices"
	"time"
)

type WeatherUseCase struct {
	vs WeatherReq
}

func NewWeatherUseCase(vs WeatherReq) *WeatherUseCase {
	return &WeatherUseCase{vs: vs}
}

var _ WeatherContract = (*WeatherUseCase)(nil)

func (w WeatherUseCase) GetWeatherInfo(dateFrom string, dateTo string, city string) (weather.ResponseData, error) {
	if city == "" {
		return weather.ResponseData{}, fmt.Errorf("location is not specified")
	}

	if dateFrom == "" {
		dateFrom = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	}

	if dateTo == "" {
		dateTo = time.Now().Format("2006-01-02")
	}

	req, err := w.vs.InitRequest(dateFrom, dateTo, city)
	if err != nil {
		return weather.ResponseData{}, err
	}

	resp, err := w.vs.SendRequest(req)
	if err != nil {
		return weather.ResponseData{}, err
	}

	parsedResult, err := w.vs.DecodeResponse(resp)
	if err != nil {
		return weather.ResponseData{}, err
	}

	var respData weather.ResponseData

	calcTemp(parsedResult, &respData)

	return respData, nil
}

func calcTemp(parsedResult *weather.WeatherData, respData *weather.ResponseData) {

	var avgVal float64
	var valuesMedian []float64

	for i, day := range parsedResult.Days {
		if i == 0 || day.Tempmin < respData.TemperatureC.Min {
			respData.TemperatureC.Min = day.Tempmin
		}
		if i == 0 || day.Tempmax > respData.TemperatureC.Max {
			respData.TemperatureC.Max = day.Tempmax
		}
		for _, hour := range day.Hours {
			avgVal += hour.Temp
			valuesMedian = append(valuesMedian, hour.Temp)
		}
	}

	respData.TemperatureC.Average = avgVal / float64(len(parsedResult.Days)*24)

	if len(valuesMedian) > 0 {
		respData.TemperatureC.Median = calculateMedian(valuesMedian)
	}
}

func calculateMedian(values []float64) float64 {
	slices.Sort(values)
	length := len(values)
	if length == 0 {
		return 0
	}
	if length%2 == 0 {
		return (values[length/2-1] + values[length/2]) / 2
	}
	return values[length/2]
}
