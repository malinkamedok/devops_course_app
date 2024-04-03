package telegram

import "devops_course_app/internal/usecase"

type TelegramBot struct {
	chatID   string
	apiToken string
}

func NewTGReq(chatID string, apiToken string) *TelegramBot {
	return &TelegramBot{chatID: chatID, apiToken: apiToken}
}

var _ usecase.TelegramReq = (*TelegramBot)(nil)
