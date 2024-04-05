package v1

import (
	gitlab2 "devops_course_app/internal/entity/gitlab"
	"devops_course_app/internal/usecase"
	"devops_course_app/pkg/web"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/render"

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
	var gw gitlab2.GitlabWebhook
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&gw)
	if err != nil {
		err := render.Render(w, r, web.ErrRender(err))
		if err != nil {
			log.Printf("Rendering error")
			return
		}
		return
	}
	if gw.ObjectAttributes.Action != "update" || len(gw.Changes.Labels.Previous) == 0 || len(gw.Changes.Labels.Current) == 0 {
		return
	}
	// disable alert for closing or opening issue with done label
	if (len(gw.Changes.Labels.Current) == 1 && gw.Changes.Labels.Current[0].Title == "Done") || (len(gw.Changes.Labels.Previous) == 1 && gw.Changes.Labels.Previous[0].Title == "Done") {
		return
	}
	data := wh.a.DecodeWebhook(&gw)
	if data.NewStatus == data.PreviousStatus {
		log.Println("issue status hasn't changed")
		return
	}
	err = wh.a.SendAlert(data)
	if err != nil {
		err := render.Render(w, r, web.ErrRender(err))
		if err != nil {
			log.Printf("Rendering error")
			return
		}
		return
	}
}
