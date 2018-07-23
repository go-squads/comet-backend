package domain

type ApplicationConfiguration struct {
	NamespaceID    int             `json:"namespaceId"`
	Version        int             `json:"version"`
	Configurations []Configuration `json:"configurations"`
}
