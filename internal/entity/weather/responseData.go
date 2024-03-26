package weather

type ResponseData struct {
	City         string `json:"city"`
	From         string `json:"from"`
	To           string `json:"to"`
	TemperatureC Stats  `json:"temperature_c"`
	Humidity     Stats  `json:"humidity"`
	PressureMb   Stats  `json:"pressure_mb"`
}
