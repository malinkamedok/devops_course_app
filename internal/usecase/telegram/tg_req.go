package telegram

import (
	"bytes"
	"devops_course_app/internal/entity/gitlab"
	"devops_course_app/internal/usecase"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	url2 "net/url"
)

type TelegramBot struct {
	chatID   string
	apiToken string
}

type ikMarkup struct {
	InlineKeyboard [][]ik `json:"inline_keyboard"`
}

type ik struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

func NewTGReq(chatID string, apiToken string) *TelegramBot {
	return &TelegramBot{chatID: chatID, apiToken: apiToken}
}

var _ usecase.TelegramReq = (*TelegramBot)(nil)

func (t TelegramBot) InitRequest(data gitlab.WebhookData) (*http.Request, error) {
	log.Println("init request")
	message := fmt.Sprintf("Issue #%d\n Student: %s\n Status: %s → %s\n", data.IssueNumber, data.StudentRepoName, data.PreviousStatus, data.NewStatus)

	var ikm ikMarkup
	ikm.InlineKeyboard[0][0].Text = "Issue"
	ikm.InlineKeyboard[0][0].Url = data.IssueURL
	ikm.InlineKeyboard[0][1].Text = "Repo"
	ikm.InlineKeyboard[0][1].Url = data.RepoURL

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(ikm)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", t.apiToken, t.chatID, url2.QueryEscape(message))

	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		log.Printf("Error in creating request")
		return nil, err
	}

	return req, nil
}

func (t TelegramBot) SendRequest(r *http.Request) error {
	log.Println("send request")
	c := http.Client{}

	resp, err := c.Do(r)
	if err != nil {
		log.Printf("Error in sending request")
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil
	}
	log.Println("response from tg: ", string(body))

	return nil
}
