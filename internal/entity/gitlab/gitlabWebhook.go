package gitlab

type GitlabWebhook struct {
	ObjectKind       string           `json:"object_kind"`
	EventType        string           `json:"event_type"`
	User             User             `json:"user"`
	Project          Project          `json:"project"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	Labels           []Label          `json:"labels"`
	Changes          Changes          `json:"changes"`
	Repository       Repository       `json:"repository"`
}
