package domain

type ConfigurationRequest struct {
	AppName   string          `json:"appName"`
	Namespace string          `json:"namespace"`
	Data      []Configuration `json:"data"`
}
