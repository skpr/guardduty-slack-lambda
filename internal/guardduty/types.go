package guardduty

type EventDetail struct {
	Severity    string `json:"severity"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
