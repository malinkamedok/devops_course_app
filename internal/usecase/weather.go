package usecase

import (
	"devops_course_app/internal/entity/weather"
	"fmt"
	"reflect"
	"sync"
)

type WeatherUseCase struct {
	vs WeatherReq
}

func NewWeatherUseCase(vs WeatherReq) *WeatherUseCase {
	return &WeatherUseCase{vs: vs}
}

var _ WeatherContract = (*WeatherUseCase)(nil)

func (w WeatherUseCase) GetWeatherInfo(dateFrom string, dateTo string, city string) (weather.ResponseData, error) {
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

	//translateFtoCtemp(parsedResult)

	var respData weather.ResponseData
	respData.City = city
	respData.From = dateFrom
	respData.To = dateTo

	calcMap := make(map[string]map[string]float64)

	var wg sync.WaitGroup
	wg.Add(3)
	go calculateValue("Temp", parsedResult, &wg, &calcMap)
	go calculateValue("Humidity", parsedResult, &wg, &calcMap)
	go calculateValue("Pressure", parsedResult, &wg, &calcMap)
	wg.Wait()

	respData.TemperatureC.Min = calcMap["Temp"]["min"]
	respData.TemperatureC.Max = calcMap["Temp"]["max"]
	respData.TemperatureC.Average = calcMap["Temp"]["average"]
	respData.TemperatureC.Median = calcMap["Temp"]["median"]

	respData.Humidity.Min = calcMap["Humidity"]["min"]
	respData.Humidity.Max = calcMap["Humidity"]["max"]
	respData.Humidity.Average = calcMap["Humidity"]["average"]
	respData.Humidity.Median = calcMap["Humidity"]["median"]

	respData.PressureMb.Min = calcMap["Pressure"]["min"]
	respData.PressureMb.Max = calcMap["Pressure"]["max"]
	respData.PressureMb.Average = calcMap["Pressure"]["average"]
	respData.PressureMb.Median = calcMap["Pressure"]["median"]

	fmt.Println("tempMin1: ", parsedResult.Days[0].Tempmin)
	fmt.Println("tempMin2: ", parsedResult.Days[1].Tempmin)

	fmt.Println("tempMAX1: ", parsedResult.Days[0].Tempmax)
	fmt.Println("tempMAX2: ", parsedResult.Days[1].Tempmax)

	return respData, nil
}

// Здесь перевод в цельсии, подсчет макс, мин и сред значений
func translateFtoCtemp(parsedResult *weather.WeatherData) {
	for i := range parsedResult.Days {
		parsedResult.Days[i].Temp = (parsedResult.Days[i].Temp - 32) * 5 / 9
		parsedResult.Days[i].Tempmin = (parsedResult.Days[i].Tempmin - 32) * 5 / 9
		parsedResult.Days[i].Tempmax = (parsedResult.Days[i].Tempmax - 32) * 5 / 9
	}
}

func calculateValue(value string, parsedResult *weather.WeatherData, wg *sync.WaitGroup, calcMap *map[string]map[string]float64) {
	defer wg.Done()

	var minVal float64
	var maxVal float64
	var avgVal float64
	var medVal float64

	// Инициализация переменных для расчета медианы
	var values []float64

	for i, day := range parsedResult.Days {
		field := reflect.ValueOf(day).FieldByName(value)
		if field.IsValid() && field.Kind() == reflect.Float64 {
			fieldValue := field.Float()
			if i == 0 || fieldValue < minVal {
				minVal = fieldValue
			}
			if i == 0 || fieldValue > maxVal {
				maxVal = fieldValue
			}
			avgVal += fieldValue
			values = append(values, fieldValue)
		}
	}

	// Расчет среднего значения
	avgVal /= float64(len(parsedResult.Days))

	// Расчет медианного значения
	if len(values) > 0 {
		medVal = calculateMedian(values)
	}

	(*calcMap)[value] = make(map[string]float64)
	(*calcMap)[value]["min"] = minVal
	(*calcMap)[value]["max"] = maxVal
	(*calcMap)[value]["avg"] = avgVal
	(*calcMap)[value]["med"] = medVal
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
