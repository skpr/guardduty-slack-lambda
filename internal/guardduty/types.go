package guardduty

type Event struct {
	Detail EventDetail `json:"detail"`
}

type EventDetail struct {
	Severity    string `json:"severity"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
