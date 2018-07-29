package domain

type ApplicationNamespace struct {
	ApplicationName string      `json:"application_name"`
	Namespace       []Namespace `json:"namespace"`
}
