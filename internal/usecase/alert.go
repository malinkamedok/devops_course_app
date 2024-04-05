package usecase

import (
	"devops_course_app/internal/entity/gitlab"
	"log"
)

type AlertUseCase struct {
	tgBot TelegramReq
}

func NewAlertUseCase(tgBot TelegramReq) *AlertUseCase {
	return &AlertUseCase{tgBot: tgBot}
}

var _ AlertContract = (*AlertUseCase)(nil)

func (a AlertUseCase) DecodeWebhook(webhook *gitlab.GitlabWebhook) *gitlab.WebhookData {
	var data gitlab.WebhookData
	data.IssueNumber = webhook.ObjectAttributes.IID
	data.StudentRepoName = webhook.Repository.Name
	if webhook.ObjectAttributes.Action == "update" {
		log.Println("update case triggered")
		log.Println(webhook.ObjectAttributes.State)
		if len(webhook.Changes.Labels.Previous) > 0 {
			data.PreviousStatus = webhook.Changes.Labels.Previous[0].Title
		}
		if len(webhook.Changes.Labels.Current) > 0 {
			data.NewStatus = webhook.Changes.Labels.Current[0].Title
		} else {
			data.NewStatus = webhook.ObjectAttributes.State
		}
	} else if webhook.ObjectAttributes.Action == "close" {
		log.Println("close case triggered")
		if len(webhook.Labels) > 1 {
			data.PreviousStatus = webhook.Labels[0].Title
			data.NewStatus = webhook.Labels[1].Title
		}
	}
	data.IssueURL = webhook.ObjectAttributes.URL
	data.RepoURL = webhook.Repository.Homepage
	return &data
}

func (a AlertUseCase) SendAlert(data *gitlab.WebhookData) error {
	req, err := a.tgBot.InitRequest(*data)
	if err != nil {
		return err
	}
	err = a.tgBot.SendRequest(req)
	if err != nil {
		return err
	}
	return nil
}
