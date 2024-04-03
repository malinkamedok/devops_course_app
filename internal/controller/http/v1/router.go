package v1

import (
	"devops_course_app/internal/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter(handler *chi.Mux, c usecase.CurrencyContract, w usecase.WeatherContract, a usecase.AlertContract) {
	handler.Route("/info", func(r chi.Router) {
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Access-Control-Allow-Origin", "X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "origin", "x-requested-with"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}))
		NewInfoRoutes(r, c, w)
	})

	handler.Route("/webhook", func(r chi.Router) {
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Access-Control-Allow-Origin", "X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "origin", "x-requested-with"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}))
		NewWebHookRoutes(r, a)
	})

	handler.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
}
