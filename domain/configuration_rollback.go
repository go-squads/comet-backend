package domain

type ConfigurationRollback struct {
	Appname       string `json:"application_name"`
	NamespaceName string `json:"namespace_name"`
	Version       int    `json:"version"`
}
