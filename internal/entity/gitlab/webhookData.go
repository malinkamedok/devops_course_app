package gitlab

type WebhookData struct {
	IssueNumber     int
	StudentRepoName string
	PreviousStatus  string
	NewStatus       string
	IssueURL        string
	RepoURL         string
}
