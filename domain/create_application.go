package domain

type CreateApplication struct {
	ApplicationsName string `json:"app_name"`
	NamespaceName    string `json:"namespaces_name"`
}
