package domain

type Configuration struct {
	NamespaceID int    `json:"namespace"`
	Version     int    `json:"version"`
	Key         string `json:"key"`
	Value       string `json:"value"`
}
