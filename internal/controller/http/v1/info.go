package v1

import (
	"devops_course_app/internal/entity/weather"
	"devops_course_app/internal/usecase"
	"devops_course_app/pkg/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

type infoRoutes struct {
	c usecase.CurrencyContract
	w usecase.WeatherContract
}

func NewInfoRoutes(routes chi.Router, c usecase.CurrencyContract, w usecase.WeatherContract) {
	ir := &infoRoutes{c: c, w: w}

	routes.Get("/currency", ir.getCurrencyRate)
	routes.Get("/weather", ir.getWeather)
}

type resp struct {
	Data    map[string]string `json:"data"`
	Service string            `json:"service"`
}

type respWeather struct {
	Data    weather.ResponseData `json:"data"`
	Service string               `json:"service"`
}

func (i *infoRoutes) getCurrencyRate(w http.ResponseWriter, r *http.Request) {
	currency := r.URL.Query().Get("currency")
	date := r.URL.Query().Get("date")

	response, err := i.c.GetCurrencyRate(currency, date)
	if err != nil {
		err := render.Render(w, r, web.ErrRender(err))
		if err != nil {
			log.Printf("Rendering error")
			return
		}
		return
	}
	responseJSON := resp{Data: map[string]string{currency: response}, Service: "currency"}
	render.JSON(w, r, responseJSON)
}

func (i *infoRoutes) getWeather(w http.ResponseWriter, r *http.Request) {
	dateFrom := r.URL.Query().Get("from")
	dateTo := r.URL.Query().Get("to")
	city := r.URL.Query().Get("city")

	response, err := i.w.GetWeatherInfo(dateFrom, dateTo, city)
	if err != nil {
		err := render.Render(w, r, web.ErrRender(err))
		if err != nil {
			log.Printf("Rendering error")
			return
		}
		return
	}
	responseJSON := respWeather{Data: response, Service: "weather"}
	render.JSON(w, r, responseJSON)
}
