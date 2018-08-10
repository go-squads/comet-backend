package domain

type ApplicationConfiguration struct {
	NamespaceID    int             `json:"namespaceId,omitempty"`
	Version        int             `json:"version"`
	Configurations []Configuration `json:"configurations"`
}
