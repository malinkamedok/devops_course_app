package weather

type Hour struct {
	Datetime       string   `json:"datetime"`
	DatetimeEpoch  int64    `json:"datetimeEpoch"`
	Temp           float64  `json:"temp"`
	Feelslike      float64  `json:"feelslike"`
	Humidity       float64  `json:"humidity"`
	Dew            float64  `json:"dew"`
	Precip         float64  `json:"precip"`
	Precipprob     float64  `json:"precipprob"`
	Snow           float64  `json:"snow"`
	Snowdepth      float64  `json:"snowdepth"`
	Preciptype     []string `json:"preciptype"`
	Windgust       float64  `json:"windgust"`
	Windspeed      float64  `json:"windspeed"`
	Winddir        float64  `json:"winddir"`
	Pressure       float64  `json:"pressure"`
	Visibility     float64  `json:"visibility"`
	Cloudcover     float64  `json:"cloudcover"`
	Solarradiation float64  `json:"solarradiation"`
	Solarenergy    float64  `json:"solarenergy"`
	Uvindex        float64  `json:"uvindex"`
	Severerisk     float64  `json:"severerisk"`
	Conditions     string   `json:"conditions"`
	Icon           string   `json:"icon"`
	Stations       []string `json:"stations"`
	Source         string   `json:"source"`
}
