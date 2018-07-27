package domain

type ConfigurationHistory struct {
	Username           string `json:"username"`
	Namespace          string `json:"namespace"`
	PredecessorVersion int    `json:"predecessor_version"`
	SuccessorVersion   int    `json:"successor_version"`
	Key                string `json:"key"`
	Value              string `json:"value"`
	CreatedAt          string `json:"created_at"`
}
