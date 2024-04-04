package types

type User struct {
	ID       string `json:"userID,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Content  string `json:"content,omitempty"`
}
