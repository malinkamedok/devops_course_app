package app

import (
	"devops_course_app/internal/config"
	v1 "devops_course_app/internal/controller/http/v1"
	"devops_course_app/internal/usecase"
	"devops_course_app/internal/usecase/cbrf"
	"devops_course_app/internal/usecase/visualcrossing"
	"devops_course_app/pkg/httpserver"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
)

func Run(cfg *config.Config) {

	c := usecase.NewCurrencyUseCase(cbrf.NewCurrencyReq())
	w := usecase.NewWeatherUseCase(visualcrossing.NewVSReq(cfg.ApiKeys))

	handler := chi.NewRouter()

	v1.NewRouter(handler, c, w)

	server := httpserver.New(handler, httpserver.Port(cfg.AppPort))
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)
	log.Printf("server started")

	select {
	case s := <-interruption:
		log.Printf("signal: " + s.String())
	case err := <-server.Notify():
		log.Printf("Notify from http server: %s\n", err)
	}

	err := server.Shutdown()
	if err != nil {
		log.Printf("Http server shutdown")
	}
}
