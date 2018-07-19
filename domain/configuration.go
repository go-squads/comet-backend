package domain

type Configuration struct {
	NamespaceID int    `json:"namespaceId"`
	Version     int    `json:"version"`
	Key         string `json:"key"`
	Value       string `json:"value"`
}
