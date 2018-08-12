package domain

type ApplicationNamespace struct {
	ApplicationName string   `json:"application_name,omitempty"`
	Namespace       []string `json:"namespace,omitempty"`
}
