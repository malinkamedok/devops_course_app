package gitlab

type ChangesLabels struct {
	Previous []Label `json:"previous"`
	Current  []Label `json:"current"`
}
