package models

type Asset struct {
	Checksum    map[string]string `json:"checksum"`
	DownloadURL string            `json:"downloadUrl"`
	Format      string            `json:"format"`
	ID          string            `json:"id"`
	Path        string            `json:"path"`
	Repository  string            `json:"repository"`
}

type AssetListResponse struct {
	Items             []*Asset `json:"items"`
	ContinuationToken string   `json:"continuationToken"`
}

type AssetFilter struct {
	// Repository from which you would like to retrieve assets.
	Repository string

	// (Optional) - A token returned by a prior request. If present, the next page of results are returned
	ContinuationToken string
}
