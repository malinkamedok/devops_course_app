package weather

type ResponseData struct {
	City         string `json:"city"`
	From         string `json:"date_from"`
	To           string `json:"date_to"`
	TemperatureC Stats  `json:"temperature_c"`
	Humidity     Stats  `json:"humidity"`
	PressureMb   Stats  `json:"pressure_mb"`
}
