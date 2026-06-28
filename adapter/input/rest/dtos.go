package rest

type MessageDTO struct {
	Id        string         `json:"id"`
	Type      string         `json:"type"`
	To        string         `json:"to"`
	Variables map[string]any `json:"variables,omitempty"`
}
