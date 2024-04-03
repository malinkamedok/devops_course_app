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
	log.Println("decode webhook")
	var data gitlab.WebhookData
	data.IssueNumber = webhook.ObjectAttributes.IID
	data.StudentRepoName = webhook.Repository.Name
	data.PreviousStatus = webhook.Changes.Labels.Previous[0].Title
	data.NewStatus = webhook.Changes.Labels.Current[0].Title
	data.IssueURL = webhook.ObjectAttributes.URL
	data.RepoURL = webhook.Repository.Homepage
	return &data
}

func (a AlertUseCase) SendAlert(data *gitlab.WebhookData) error {
	log.Println("send webhook")
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
