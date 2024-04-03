package v1

import (
	"devops_course_app/internal/usecase"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type webhookRoutes struct {
	a usecase.AlertContract
}

func NewWebHookRoutes(routes chi.Router, a usecase.AlertContract) {
	wr := &webhookRoutes{a: a}

	routes.Post("/gitlab", wr.gitlabWebhook)
}

func (wh *webhookRoutes) gitlabWebhook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Reading body error")
	}
	fmt.Println(string(body))
}
