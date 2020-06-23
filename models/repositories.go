package models

type RepositoryV1 struct {
	Attributes map[string]interface{} `json:"attributes"`
	Format     string                 `json:"format"`
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	URL        string                 `json:"url"`
}
