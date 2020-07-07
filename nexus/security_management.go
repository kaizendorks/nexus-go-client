package nexus

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

type SecurityManagementAPI api

// ListUserSources retrieves a list of the available user sources.
//	api endpoint: GET /beta/security/user-sources
//	responses:
// 		200: successful operation returns UserSource slice"
// 		403: The user does not have permission to perform the operation.
func (a SecurityManagementAPI) List() ([]UserSource, error) {
	us := []UserSource{}
	path := fmt.Sprintf("beta/security/user-sources")

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return us, err
	}
	err = json.Unmarshal(resp, &us)
	return us, err
}
