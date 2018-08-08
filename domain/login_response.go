package domain

type LoginResponse struct {
	Status    int    `json:"status"`
	Fullname  string `json:"full_name,omitempty"`
	RoleBased string `json:"user_role,omitempty"`
	Message   string `json:"message"`
	Token     string `json:"token,omitempty"`
}
