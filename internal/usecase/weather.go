package usecase

type WeatherUseCase struct {
	vs WeatherReq
}

func NewWeatherUseCase(vs WeatherReq) *WeatherUseCase {
	return &WeatherUseCase{vs: vs}
}

func (w WeatherUseCase) GetWeatherInfo(city string) (string, error) {
	return "aboba", nil
}

var _ WeatherContract = (*WeatherUseCase)(nil)
