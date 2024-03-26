package weather

type WeatherData struct {
	QueryCost       int     `json:"queryCost"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	ResolvedAddress string  `json:"resolvedAddress"`
	Address         string  `json:"address"`
	Timezone        string  `json:"timezone"`
	Tzoffset        float64 `json:"tzoffset"`
	Description     string  `json:"description"`
	Days            []Day   `json:"days"`
}
