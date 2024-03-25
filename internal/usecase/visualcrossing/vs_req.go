package visualcrossing

import "devops_course_app/internal/usecase"

type WeatherVS struct{}

func NewVSReq() *WeatherVS {
	return &WeatherVS{}
}

var _ usecase.WeatherReq = (*WeatherVS)(nil)
