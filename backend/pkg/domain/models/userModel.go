package models

type User struct {
	ID       string `json:"_id"`
	Type     string `json:"_type"`
	Username string `json:"username"`
	Password string `json:"password"`
}
