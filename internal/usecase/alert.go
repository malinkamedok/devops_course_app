package usecase

import (
	"devops_course_app/internal/entity/gitlab"
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
	data.WhoChanged = webhook.User.Username

	if len(webhook.Changes.Labels.Previous) > 0 {
		for i, label := range webhook.Changes.Labels.Previous {
			if i > 0 {
				data.PreviousStatus = data.PreviousStatus + ", "
			}
			data.PreviousStatus = data.PreviousStatus + label.Title
		}
	} else {
		data.PreviousStatus = webhook.ObjectAttributes.State
	}
	if len(webhook.Changes.Labels.Current) > 0 {
		for i, label := range webhook.Changes.Labels.Current {
			if i > 0 {
				data.NewStatus = data.NewStatus + ", "
			}
			data.NewStatus = data.NewStatus + label.Title
		}
	} else {
		data.NewStatus = webhook.ObjectAttributes.State
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
