package visualcrossing

import (
	"devops_course_app/internal/entity/weather"
	"devops_course_app/internal/usecase"
	"encoding/json"
	"log"
	"net/http"
)

type WeatherVS struct {
	apiKey string
}

func NewVSReq(apiKey string) *WeatherVS {
	return &WeatherVS{apiKey: apiKey}
}

var _ usecase.WeatherReq = (*WeatherVS)(nil)

func (w WeatherVS) InitRequest(dateFrom string, dateTo string, city string) (*http.Request, error) {
	url := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/" + city + "/" + dateFrom + "/" + dateTo + "?unitGroup=metric&key=" + w.apiKey

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error in creating request")
		return nil, err
	}

	return req, nil
}

func (w WeatherVS) SendRequest(r *http.Request) (*http.Response, error) {
	c := http.Client{}

	resp, err := c.Do(r)
	if err != nil {
		log.Printf("Error in sending request")
		return nil, err
	}

	return resp, nil
}

func (w WeatherVS) DecodeResponse(response *http.Response) (*weather.WeatherData, error) {
	defer response.Body.Close()

	var resp weather.WeatherData
	err := json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		log.Printf("Error in decoding response")
		return nil, err
	}

	return &resp, nil
}
