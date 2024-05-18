package ui

type blockJsonObject struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Content    string            `json:"content"`
	Parameters map[string]string `json:"parameters"`
	Children   []blockJsonObject `json:"children"`
}
