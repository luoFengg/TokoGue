package web

type UserLoginResponse struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token,omitempty"`
}