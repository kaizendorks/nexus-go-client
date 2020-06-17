package nexus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AssetsAPI api

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

// List assets
// api endpoint: GET /v1​/assets
// parameters:
// 		continuationToken
// 			description: A token returned by a prior request. If present, the next page of results are returned
// 			required: false
// 		repository
// 			description: Repository from which you would like to retrieve assets.
// 			required: true
// responses:
// 		200: Successful operation returns AssetListResponse and nil error
// 		403: Insufficient permissions to list assets
// 		422: Parameter 'repository' is required (this methods should never get into this state.)
func (a AssetsAPI) List(repository, continuationToken string) (*AssetListResponse, error) {
	assetResp := &AssetListResponse{}
	path := fmt.Sprintf("v1/assets?repository=%s", repository)
	if continuationToken != "" {
		path = fmt.Sprintf("%s&%s", path, repository)
	}

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return assetResp, err
	}
	err = json.Unmarshal(resp, assetResp)
	return assetResp, err
}

// Get a single asset
// endpoint: GET /v1​/assets​/{id}
// parameters:
// 		assetId:
//    	description: ID of the asset to get
//    	required: true
// 	responses:
// 		200: Successful operation returns Asset and nil error
// 		403: Insufficient permissions to get asset
// 		404: Asset not found
// 		422: Malformed ID
func (a AssetsAPI) Get(assetId string) (*Asset, error) {
	asset := &Asset{}
	path := fmt.Sprintf("v1/assets/%s", assetId)

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return asset, err
	}

	err = json.Unmarshal(resp, asset)
	return asset, err
}

// Delete a single asset
// endpoint: DELETE /v1​/assets​/{id}
// parameters:
// 		assetId:
//    	description: ID of the asset to delete
//    	required: true
// responses:
// 		204: Asset was successfully deleted
// 		403: Insufficient permissions to delete asset
// 		404: Asset not found
// 		422: Malformed ID
func (a AssetsAPI) Delete(assetId string) error {
	path := fmt.Sprintf("v1/assets/%s", assetId)

	_, err := a.client.sendRequest(http.MethodDelete, path, nil, nil)
	return err
}
