package usecase

import (
	"devops_course_app/internal/entity"
	"net/http"
)

type (
	CurrencyReq interface {
		InitRequest(dateFormatted string) (*http.Request, error)
		SendRequest(r *http.Request) (*http.Response, error)
		DecodeResponse(response *http.Response) (*entity.ValCurs, error)
		FindCurrencyRate(currency string, currencyRates *entity.ValCurs) (string, error)
	}

	CurrencyContract interface {
		GetCurrencyRate(currency string, date string) (string, error)
	}

	WeatherReq interface {
	}

	WeatherContract interface {
		GetWeatherInfo(city string) (string, error)
	}
)
