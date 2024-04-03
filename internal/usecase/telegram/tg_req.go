package telegram

import (
	"devops_course_app/internal/entity/gitlab"
	"devops_course_app/internal/usecase"
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

func NewTGReq(chatID string, apiToken string) *TelegramBot {
	return &TelegramBot{chatID: chatID, apiToken: apiToken}
}

var _ usecase.TelegramReq = (*TelegramBot)(nil)

func (t TelegramBot) InitRequest(data gitlab.WebhookData) (*http.Request, error) {
	log.Println("init request")
	message := fmt.Sprintf("Issue #%d\n Student: %s\n Status: %s → %s\n", data.IssueNumber, data.StudentRepoName, data.PreviousStatus, data.NewStatus)
	keyboard := fmt.Sprintf(`{
 		"inline_keyboard": [
			[
			  {"text": "Issue", "url": %s},
			  {"text": "Repo", "url": %s}
			]
	 	]
	}`, data.IssueURL, data.RepoURL)

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s&reply_markup=%s", t.apiToken, t.chatID, url2.QueryEscape(message), url2.QueryEscape(keyboard))

	req, err := http.NewRequest("GET", url, nil)
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
