package models

type Message struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}
