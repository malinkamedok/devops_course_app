package gitlab

type ObjectAttributes struct {
	AuthorID                  int           `json:"author_id"`
	ClosedAt                  string        `json:"closed_at"`
	Confidential              bool          `json:"confidential"`
	CreatedAt                 string        `json:"created_at"`
	Description               string        `json:"description"`
	DiscussionLocked          bool          `json:"discussion_locked"`
	DueDate                   string        `json:"due_date"`
	ID                        int           `json:"id"`
	IID                       int           `json:"iid"`
	LastEditedAt              string        `json:"last_edited_at"`
	LastEditedByID            int           `json:"last_edited_by_id"`
	MilestoneID               int           `json:"milestone_id"`
	MovedToID                 int           `json:"moved_to_id"`
	DuplicatedToID            int           `json:"duplicated_to_id"`
	ProjectID                 int           `json:"project_id"`
	RelativePosition          int           `json:"relative_position"`
	StateID                   int           `json:"state_id"`
	TimeEstimate              int           `json:"time_estimate"`
	Title                     string        `json:"title"`
	UpdatedAt                 string        `json:"updated_at"`
	UpdatedByID               int           `json:"updated_by_id"`
	URL                       string        `json:"url"`
	TotalTimeSpent            int           `json:"total_time_spent"`
	TimeChange                int           `json:"time_change"`
	HumanTotalTimeSpent       string        `json:"human_total_time_spent"`
	HumanTimeChange           string        `json:"human_time_change"`
	HumanTimeEstimate         string        `json:"human_time_estimate"`
	AssigneeIDs               []int         `json:"assignee_ids"`
	AssigneeID                int           `json:"assignee_id"`
	Labels                    []Label       `json:"labels"`
	State                     string        `json:"state"`
	Severity                  string        `json:"severity"`
	CustomerRelationsContacts []interface{} `json:"customer_relations_contacts"`
	Action                    string        `json:"action"`
}
