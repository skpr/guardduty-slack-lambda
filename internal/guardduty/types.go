package guardduty

// Event received by GuardDuty.
type Event struct {
	Detail EventDetail `json:"detail"`
}

// EventDetail provided by GuardDuty.
type EventDetail struct {
	Severity    int    `json:"severity"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
