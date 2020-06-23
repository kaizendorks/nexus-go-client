package models

type Script struct {
	Content string `json:"content,omitempty"`

	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name,omitempty"`

	Type string `json:"type,omitempty"`
}

type ScriptResult struct {
	Name   string `json:"name,omitempty"`
	Result string `json:"result,omitempty"`
}
