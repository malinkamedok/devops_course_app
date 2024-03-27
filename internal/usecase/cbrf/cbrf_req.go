package cbrf

import (
	"bytes"
	"devops_course_app/internal/entity/currency"
	"devops_course_app/internal/usecase"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type CurrencyCBRF struct{}

func NewCurrencyReq() *CurrencyCBRF {
	return &CurrencyCBRF{}
}

func (i CurrencyCBRF) InitRequest(dateFormatted string) (*http.Request, error) {
	url := "https://www.cbr.ru/scripts/XML_daily.asp?date_req=" + dateFormatted

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error in creating request")
		return nil, err
	}

	// Getting 403 Forbidden error without setting this header
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")

	return req, nil
}

func (i CurrencyCBRF) SendRequest(r *http.Request) (*http.Response, error) {
	c := http.Client{}

	resp, err := c.Do(r)
	if err != nil {
		log.Printf("Error in sending request")
		return nil, err
	}

	return resp, nil
}

func (i CurrencyCBRF) DecodeResponse(response *http.Response) (*currency.ValCurs, error) {
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error in reading response")
		return nil, err
	}

	reader := bytes.NewReader(responseData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	rates := new(currency.ValCurs)

	err = decoder.Decode(rates)
	if err != nil {
		log.Printf("Error in decoding response")
		return nil, err
	}

	return rates, nil
}

func (i CurrencyCBRF) FindCurrencyRate(currency string, currencyRates *currency.ValCurs) (float64, error) {
	for _, v := range currencyRates.Valutes {
		if v.CharCode == currency {
			return strconv.ParseFloat(strings.Replace(v.Value, ",", ".", 1), 64)
		}
	}
	return 0, fmt.Errorf("Currency or rate not found")
}

var _ usecase.CurrencyReq = (*CurrencyCBRF)(nil)
