package usecase

import (
	"devops_course_app/internal/entity/weather"
	"fmt"
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
		//city = "SaintPetersburg"
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
	respData.City = city
	respData.From = dateFrom
	respData.To = dateTo

	calcTemp(parsedResult, &respData)
	calcHumidity(parsedResult, &respData)
	calcPressure(parsedResult, &respData)

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
		avgVal += day.Temp
		valuesMedian = append(valuesMedian, day.Temp)
	}

	respData.TemperatureC.Average = avgVal / float64(len(parsedResult.Days))

	if len(valuesMedian) > 0 {
		respData.TemperatureC.Median = calculateMedian(valuesMedian)
	}
}

func calcHumidity(parsedResult *weather.WeatherData, respData *weather.ResponseData) {
	var avgVal float64
	var valuesMedian []float64
	respData.Humidity.Min = 100
	respData.Humidity.Max = 0

	for _, day := range parsedResult.Days {
		for _, hour := range day.Hours {
			if hour.Humidity < respData.Humidity.Min {
				respData.Humidity.Min = hour.Humidity
			}
			if hour.Humidity > respData.Humidity.Max {
				respData.Humidity.Max = hour.Humidity
			}
			avgVal += hour.Humidity
			valuesMedian = append(valuesMedian, hour.Humidity)
		}
	}

	respData.Humidity.Average = avgVal / float64(len(parsedResult.Days)*24)

	if len(valuesMedian) > 0 {
		respData.Humidity.Median = calculateMedian(valuesMedian)
	}
}

func calcPressure(parsedResult *weather.WeatherData, respData *weather.ResponseData) {
	var avgVal float64
	var valuesMedian []float64
	respData.PressureMb.Min = 2000
	respData.PressureMb.Max = 0

	for _, day := range parsedResult.Days {
		for _, hour := range day.Hours {
			if hour.Pressure < respData.PressureMb.Min {
				respData.PressureMb.Min = hour.Pressure
			}
			if hour.Pressure > respData.PressureMb.Max {
				respData.PressureMb.Max = hour.Pressure
			}
			avgVal += hour.Pressure
			valuesMedian = append(valuesMedian, hour.Pressure)
		}
	}

	respData.PressureMb.Average = avgVal / float64(len(parsedResult.Days)*24)

	if len(valuesMedian) > 0 {
		respData.PressureMb.Median = calculateMedian(valuesMedian)
	}
}

func calculateMedian(values []float64) float64 {
	length := len(values)
	if length == 0 {
		return 0
	}
	if length%2 == 0 {
		return (values[length/2-1] + values[length/2]) / 2
	}
	return values[length/2]
}
