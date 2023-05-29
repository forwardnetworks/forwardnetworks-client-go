package forwardnetworks

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetExternalId function retrieves the external ID of the specified network.
func (c *Client) GetExternalId(networkId string) (*ExternalId, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/networks/%s/cloudAccounts/aws/assumeRole/externalId", c.HostURL, networkId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	externalId := ExternalId{}
	err = json.Unmarshal(body, &externalId)
	if err != nil {
		return nil, err
	}

	return &externalId, nil
}