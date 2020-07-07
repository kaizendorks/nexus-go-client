package nexus

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

type AssetsAPI api

// List assets
//	api endpoint: GET /v1/assets
//	parameters:
// 		af: AssetFilter object consisting of options to filter the results by.
//	responses:
// 		200: Successful operation returns AssetListResponse and nil error
// 		403: Insufficient permissions to list assets
// 		422: Parameter 'repository' is required (this method should never get into this state when using the go client.)
func (a AssetsAPI) List(af AssetFilter) (AssetListResponse, error) {
	path := fmt.Sprintf("v1/assets?repository=%s", af.Repository)
	if af.ContinuationToken != "" {
		path = fmt.Sprintf("%s&%s", path, af.ContinuationToken)
	}

	assetResp := AssetListResponse{}

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return assetResp, err
	}
	err = json.Unmarshal(resp, &assetResp)
	return assetResp, err
}

// Get a single asset
//	endpoint: GET /v1/assets/{id}
//	parameters:
//		assetId: ID of the asset to get
//	responses:
// 		200: Successful operation returns Asset and nil error
// 		403: Insufficient permissions to get asset
// 		404: Asset not found
// 		422: Malformed ID
func (a AssetsAPI) Get(assetId string) (Asset, error) {
	path := fmt.Sprintf("v1/assets/%s", assetId)
	asset := Asset{}

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return asset, err
	}

	err = json.Unmarshal(resp, &asset)
	return asset, err
}

// Delete a single asset
//	endpoint: DELETE /v1/assets/{id}
// 	parameters:
// 		assetId: ID of the asset to delete
// 	responses:
// 		204: Asset was successfully deleted
// 		403: Insufficient permissions to delete asset
// 		404: Asset not found
// 		422: Malformed ID
func (a AssetsAPI) Delete(assetId string) error {
	path := fmt.Sprintf("v1/assets/%s", assetId)

	_, err := a.client.sendRequest(http.MethodDelete, path, nil, nil)
	return err
}
