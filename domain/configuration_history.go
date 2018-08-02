package domain

type ConfigurationHistory struct {
	Username           string          `json:"username"`
	Namespace          string          `json:"namespace"`
	PredecessorVersion int             `json:"predecessor_version"`
	SuccessorVersion   int             `json:"successor_version"`
	CreatedAt          string          `json:"created_at"`
	Deleted            []Configuration `json:"deleted,omitempty"`
	Changed            []Configuration `json:"changed,omitempty"`
	Created            []Configuration `json:"created,omitempty"`
}
