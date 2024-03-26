package usecase

import (
	"devops_course_app/internal/entity/currency"
	"devops_course_app/internal/entity/weather"
	"net/http"
)

type (
	CurrencyReq interface {
		InitRequest(dateFormatted string) (*http.Request, error)
		SendRequest(r *http.Request) (*http.Response, error)
		DecodeResponse(response *http.Response) (*currency.ValCurs, error)
		FindCurrencyRate(currency string, currencyRates *currency.ValCurs) (string, error)
	}

	CurrencyContract interface {
		GetCurrencyRate(currency string, date string) (string, error)
	}

	WeatherReq interface {
		InitRequest(dateFrom string, dateTo string, city string) (*http.Request, error)
		SendRequest(r *http.Request) (*http.Response, error)
		DecodeResponse(response *http.Response) (*weather.WeatherData, error)
	}

	WeatherContract interface {
		GetWeatherInfo(dateFrom string, dateTo string, city string) (weather.ResponseData, error)
	}
)
