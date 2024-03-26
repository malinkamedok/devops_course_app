package weather

type Day struct {
	Datetime       string   `json:"datetime"`
	DatetimeEpoch  int64    `json:"datetimeEpoch"`
	Tempmax        float64  `json:"tempmax"`
	Tempmin        float64  `json:"tempmin"`
	Temp           float64  `json:"temp"`
	Feelslikemax   float64  `json:"feelslikemax"`
	Feelslikemin   float64  `json:"feelslikemin"`
	Feelslike      float64  `json:"feelslike"`
	Dew            float64  `json:"dew"`
	Humidity       float64  `json:"humidity"`
	Precip         float64  `json:"precip"`
	Precipprob     float64  `json:"precipprob"`
	Precipcover    float64  `json:"precipcover"`
	Preciptype     []string `json:"preciptype"`
	Snow           float64  `json:"snow"`
	Snowdepth      float64  `json:"snowdepth"`
	Windgust       float64  `json:"windgust"`
	Windspeed      float64  `json:"windspeed"`
	Winddir        float64  `json:"winddir"`
	Pressure       float64  `json:"pressure"`
	Cloudcover     float64  `json:"cloudcover"`
	Visibility     float64  `json:"visibility"`
	Solarradiation float64  `json:"solarradiation"`
	Solarenergy    float64  `json:"solarenergy"`
	Uvindex        float64  `json:"uvindex"`
	Severerisk     float64  `json:"severerisk"`
	Sunrise        string   `json:"sunrise"`
	SunriseEpoch   int64    `json:"sunriseEpoch"`
	Sunset         string   `json:"sunset"`
	SunsetEpoch    int64    `json:"sunsetEpoch"`
	Moonphase      float64  `json:"moonphase"`
	Description    string   `json:"description"`
	Icon           string   `json:"icon"`
	Stations       []string `json:"stations"`
	Source         string   `json:"source"`
	Hours          []Hour   `json:"hours"`
}
